package targets

import (
	"bytes"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/openkit-devtools/openkit/internal/syncer"
)

func desiredGeminiCommands(base fs.FS) ([]syncer.DesiredFile, error) {
	entries, err := fs.ReadDir(base, "base/commands")
	if err != nil {
		return nil, err
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	var out []syncer.DesiredFile
	for _, ent := range entries {
		if ent.IsDir() {
			continue
		}
		name := ent.Name()
		if name == "README.md" {
			continue
		}
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		embeddedPath := path.Join("base/commands", name)
		b, err := fs.ReadFile(base, embeddedPath)
		if err != nil {
			return nil, err
		}

		desc, body := parseMarkdownFrontmatter(b)
		prompt := strings.ReplaceAll(string(body), "$ARGUMENTS", "{{args}}")
		cmdID := strings.TrimSuffix(name, ".md")

		tomlBytes, err := geminiCommandToml(desc, prompt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}

		out = append(out, syncer.DesiredFile{
			OutputPath: ".gemini/commands/openkit/" + cmdID + ".toml",
			Bytes:      tomlBytes,
			ArtifactID: "generated/gemini/commands/openkit/" + cmdID + ".toml",
			Mode:       "copy",
		})
	}

	sort.Slice(out, func(i, j int) bool { return out[i].OutputPath < out[j].OutputPath })
	return out, nil
}

func parseMarkdownFrontmatter(b []byte) (description string, body []byte) {
	// Strip UTF-8 BOM if present.
	b = bytes.TrimPrefix(b, []byte{0xEF, 0xBB, 0xBF})
	if !bytes.HasPrefix(b, []byte("---\n")) {
		return "", b
	}

	end := bytes.Index(b[len("---\n"):], []byte("\n---\n"))
	if end == -1 {
		return "", b
	}

	fmStart := len("---\n")
	fmEnd := fmStart + end
	fm := b[fmStart:fmEnd]
	body = b[fmEnd+len("\n---\n"):]

	for _, line := range strings.Split(string(fm), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.HasPrefix(line, "description:") {
			continue
		}
		v := strings.TrimSpace(strings.TrimPrefix(line, "description:"))
		v = strings.Trim(v, "\"'")
		return v, body
	}

	return "", body
}

func geminiCommandToml(description string, prompt string) ([]byte, error) {
	var buf bytes.Buffer

	if strings.TrimSpace(description) != "" {
		buf.WriteString("description = ")
		buf.WriteString(tomlString(description))
		buf.WriteString("\n")
	}

	buf.WriteString("prompt = ")
	buf.WriteString(tomlMultilineOrBasic(prompt))
	buf.WriteString("\n")

	return buf.Bytes(), nil
}

func tomlMultilineOrBasic(s string) string {
	if strings.Contains(s, "\"\"\"") {
		return tomlString(s)
	}
	return "\"\"\"\n" + s + "\"\"\""
}

func tomlString(s string) string {
	// A basic TOML string with escapes.
	// This is used as a fallback for content that cannot fit a multiline string.
	var b strings.Builder
	b.Grow(len(s) + 2)
	b.WriteByte('"')
	for _, r := range s {
		switch r {
		case '\\':
			b.WriteString("\\\\")
		case '"':
			b.WriteString("\\\"")
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
	return b.String()
}
