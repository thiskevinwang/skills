---
name: setup-pstack
description: Configure optional pstack model choices per role in Codex. Use for $setup-pstack or requests to change pstack delegate models.
---

# Setup pstack

Write `${CODEX_HOME:-$HOME/.codex}/pstack-models.md`. The poteto-mode Codex adapter reads this optional file before delegation. It is not an always-applied rule and does not change global Codex settings.

## Detect and read

Inspect the model IDs and reasoning levels accepted by this session's subagent tools. Those schemas are the source for choices. Do not translate Cursor model IDs into guessed Codex IDs. If no explicit model choices are exposed, use parent inheritance and state that model selection is unavailable.

Read the existing role file before editing it. Preserve roles the user did not ask to change. Without a file, start every role on `inherit-parent`.

## Choose

Apply model choices the user already gave. Otherwise show the current role choices and ask which to change. Offer only confirmed IDs, `inherit-parent`, and `auto`. The aliases mean omit model arguments. They are never sent as literal model IDs.

Panel roles take a comma-separated list. Each entry counts as one reviewer, including repeated inheritance entries. A panel on one family provides separate reviews, not evidence across model families. The arena cross-judge pool is a selection pool, not a reviewer count.

Use these role labels when needed:

- `feature, refactoring`
- `bug-fix`
- `perf-issue`
- `hillclimb`
- `judgment and prose`
- `hardest tasks`
- `how explorer`, `how explainer`, `how critics`
- `why investigators`, `why synthesizer`
- `reflect tooling`, `reflect judgment, divergent, synthesizer`
- `arena runners`, `arena cross-judge pool`
- `swarm workers`
- `architect runners`
- `interrogate reviewers`

## Validate and write

Validate each real model ID against the current tool schema. Use a separate `reasoning effort for <role>` line only when the user chose a supported level. Do not encode effort or speed in a model ID. If the tool requires fresh context for an override, give that context the complete brief.

Write the role file after choices are known. Retain unmodified roles. Example:

```text
# pstack model choices for Codex. Missing roles inherit the parent.
feature, refactoring: inherit-parent
judgment and prose: inherit-parent
how critics: inherit-parent, inherit-parent
swarm workers: inherit-parent
```

Report the path and changed roles. The next poteto-mode delegation reads the file. This does not make Cursor-only sibling skills native Codex workflows outside the poteto-mode adapter.
