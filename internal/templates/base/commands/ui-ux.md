---
description: Create design systems, UI components, and UX audits with psychology principles
subtask: true
---

**UI/UX DESIGN MODE ACTIVATED**

**User task:** $ARGUMENTS

**If $ARGUMENTS is empty:**
Use the question tool to ask:
```javascript
question({
  questions: [{
      question: "What would you like to do?",
      header: "UI/UX Task",
      options: [
        { label: "Design system", description: "Create a full design system" },
        { label: "UI/component", description: "Design a specific UI or component" },
        { label: "UX audit", description: "Run a UX audit of the project" }
      ]
    }]
})
```

**Capabilities:**
-  **Design System:** Color palette, typography, spacing, base components
-  **UI Design:** Specific layouts, pages, components
-  **UX Audit:** Analysis against psychology principles

---

## Option 1: Create a Design System

**Execution prompt:**
"Crie um design system completo para o projeto atual incluindo:

### 1. Color Palette
- Primary colors (brand)
- Secondary colors
- Neutral grays
- Semantic colors (success, warning, error, info)
- Verifique contrast ratios WCAG AA

### 2. Typography Scale
- Font families
- Heading scale (H1-H6)
- Body text sizes
- Line heights
- Font weights

### 3. Spacing System
- Base unit (4px)
- Scale: xs, sm, md, lg, xl, 2xl, etc.
- Usage guidelines

### 4. Component Library
- Buttons (variants, states)
- Inputs (text, select, checkbox, radio)
- Cards
- Modals/Dialogs
- Navigation
- Feedback (alerts, toasts)

### 5. Design Tokens
Export as:
- Tailwind config
- CSS variables
- JSON tokens

**Output:** `design-system.md` + config files"

---

## Option 2: Component/UI Design

**Execution prompt:**
"Design o componente/interface solicitada: $ARGUMENTS

Inclua:
1. **Layout & Structure**
   - Wireframe/prototype
   - Responsive breakpoints
   - Grid/flex layout

2. **Visual Design**
   - Cores aplicadas
   - Tipografia
   - Espaçamento
   - Sombras/elevation

3. **Interaction States**
   - Default
   - Hover
   - Active/Pressed
   - Focus
   - Disabled
   - Loading
   - Error

4. **Accessibility**
   - ARIA labels
   - Keyboard navigation
   - Screen reader support
   - Color contrast

5. **Implementation Notes**
   - Tech stack (React, Vue, etc.)
   - Props/API
   - Exemplo de uso

**Use skill `frontend-design` and agent `frontend-specialist`.**"

---

## Option 3: UX Audit

**Execution prompt:**
"Execute UX audit completo do projeto atual.

### Psychology Laws Checklist:

**Hick's Law** (Minimize choices)
- [ ] Tarefas críticas têm poucas opções
- [ ] Complexidade dividida em steps
- [ ] Progressive disclosure aplicada

**Fitts's Law** (Target sizes)
- [ ] Touch targets >= 44x44px
- [ ] Ações importantes são maiores/mais próximas
- [ ] Consideração para mobile

**Jakob's Law** (Familiar patterns)
- [ ] Usa padrões familiares
- [ ] Não reinventa navegação
- [ ] Consistente com convenções web

**Law of Proximity** (Grouping)
- [ ] Elementos relacionados agrupados
- [ ] Spacing indica relacionamento
- [ ] Hierarquia visual clara

### Visual Design Audit:

**Hierarchy**
- [ ] Focal points claros
- [ ] Elementos importantes proeminentes
- [ ] Scanning patterns (Z, F)

**Contrast**
- [ ] Texto legível
- [ ] WCAG AA compliance (4.5:1)
- [ ] Elementos importantes destacam-se

**Consistency**
- [ ] Mesmos elementos parecem iguais
- [ ] Espaçamento consistente
- [ ] Comportamento previsível

### Trust & Credibility:
- [ ] Aparência profissional
- [ ] Value proposition clara
- [ ] Social proof (se aplicável)
- [ ] Contato/suporte visível
- [ ] Indicadores de segurança

**Output:** Audit report with score and prioritized action items."

---

## Available Tools

```bash
# Automated UX audit
!`python .opencode/skills/frontend-design/scripts/ux_audit.py .`

# Accessibility check
!`python .opencode/skills/frontend-design/scripts/accessibility_checker.py .`

# Design system generator (optional)
!`python .opencode/skills/frontend-design/scripts/design_system.py`
```

---

## Design Documentation

Sempre crie documentação:
```markdown
# Design: [Name]

## Overview
[Context and goals]

## Design Decisions
[Why we chose X over Y]

## Specs
[Colors, typography, spacing]

## States
[All component states]

## Usage
[Code examples]

## Accessibility
[A11y considerations]
```

---

** STOP POINT:**
> Use the question tool to ask "Design complete. Review before proceeding to implementation?"

**Next step:** Use `/impl` to implement the design.
