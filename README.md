# Agent Skills

This repository stores installable agent skills. Each skill lives under
`skills/<skill-name>` with its instructions and any supporting metadata or
resources kept together.

## Installing

```sh
npx skills add thiskevinwang/skills
# or
gh skill install thiskevinwang/skills
```

## Available skills

- [go-callgraph](skills/go-callgraph/SKILL.md): Build and display Go call graphs.
- [create-quiz](skills/create-quiz/SKILL.md): Create a quiz and answer key with source citations.
- [quiz-me](skills/quiz-me/SKILL.md): Learn concepts with one question at a time.

### Principles

These 24 skills use the latest pstack instructions with Codex runtime adaptations. In Codex,
invoke them explicitly with `$principle-<name>`.

- [principle-boundary-discipline](skills/principle-boundary-discipline/SKILL.md)
- [principle-build-the-lever](skills/principle-build-the-lever/SKILL.md)
- [principle-encode-lessons-in-structure](skills/principle-encode-lessons-in-structure/SKILL.md)
- [principle-exhaust-the-design-space](skills/principle-exhaust-the-design-space/SKILL.md)
- [principle-experience-first](skills/principle-experience-first/SKILL.md)
- [principle-fix-root-causes](skills/principle-fix-root-causes/SKILL.md)
- [principle-foundational-thinking](skills/principle-foundational-thinking/SKILL.md)
- [principle-guard-the-context-window](skills/principle-guard-the-context-window/SKILL.md)
- [principle-laziness-protocol](skills/principle-laziness-protocol/SKILL.md)
- [principle-make-operations-idempotent](skills/principle-make-operations-idempotent/SKILL.md)
- [principle-migrate-callers-then-delete-legacy-apis](skills/principle-migrate-callers-then-delete-legacy-apis/SKILL.md)
- [principle-minimize-reader-load](skills/principle-minimize-reader-load/SKILL.md)
- [principle-model-the-domain](skills/principle-model-the-domain/SKILL.md)
- [principle-never-block-on-the-human](skills/principle-never-block-on-the-human/SKILL.md)
- [principle-outcome-oriented-execution](skills/principle-outcome-oriented-execution/SKILL.md)
- [principle-prove-it-works](skills/principle-prove-it-works/SKILL.md)
- [principle-redesign-from-first-principles](skills/principle-redesign-from-first-principles/SKILL.md)
- [principle-separate-before-serializing-shared-state](skills/principle-separate-before-serializing-shared-state/SKILL.md)
- [principle-sequence-verifiable-units](skills/principle-sequence-verifiable-units/SKILL.md)
- [principle-subtract-before-you-add](skills/principle-subtract-before-you-add/SKILL.md)
- [principle-type-system-discipline](skills/principle-type-system-discipline/SKILL.md)
- [principle-attack-the-premise](skills/principle-attack-the-premise/SKILL.md)
- [principle-explain-the-number](skills/principle-explain-the-number/SKILL.md)
- [principle-test-behavior-not-implementation](skills/principle-test-behavior-not-implementation/SKILL.md)

### Cursor plugin ports

These installed ports come from [pstack](https://github.com/cursor/plugins/tree/main/pstack),
except [thermo-nuclear-code-quality-review](skills/thermo-nuclear-code-quality-review/SKILL.md),
which comes from [thermos](https://github.com/cursor/plugins/tree/main/thermos).

- [architect](skills/architect/SKILL.md)
- [arena](skills/arena/SKILL.md)
- [automate-me](skills/automate-me/SKILL.md)
- [benchmark-checklist](skills/benchmark-checklist/SKILL.md)
- [blast-radius](skills/blast-radius/SKILL.md)
- [bro](skills/bro/SKILL.md)
- [create-verification-skill](skills/create-verification-skill/SKILL.md)
- [figure-it-out](skills/figure-it-out/SKILL.md)
- [how](skills/how/SKILL.md)
- [interrogate](skills/interrogate/SKILL.md)
- [maintain-verification-skill](skills/maintain-verification-skill/SKILL.md)
- [make-bot-ui](skills/make-bot-ui/SKILL.md)
- [no-comments](skills/no-comments/SKILL.md)
- [poteto-mode](skills/poteto-mode/SKILL.md)
- [recall](skills/recall/SKILL.md)
- [reflect](skills/reflect/SKILL.md)
- [setup-pstack](skills/setup-pstack/SKILL.md)
- [show-me-your-work](skills/show-me-your-work/SKILL.md)
- [statechart](skills/statechart/SKILL.md)
- [swarm](skills/swarm/SKILL.md)
- [tdd](skills/tdd/SKILL.md)
- [teach](skills/teach/SKILL.md)
- [technical-writing](skills/technical-writing/SKILL.md)
- [typescript-best-practices](skills/typescript-best-practices/SKILL.md)
- [unslop](skills/unslop/SKILL.md)
- [why](skills/why/SKILL.md)
- [thermo-nuclear-code-quality-review](skills/thermo-nuclear-code-quality-review/SKILL.md)

The principle skills also come from pstack. All imports preserve local Codex
adaptations and supporting files. Each imported skill includes its upstream
MIT license: pstack, copyright 2026 Lauren Tan; thermos, copyright 2026 Cursor.

The 2026-10-29 refresh imports `poteto-mode`, all 24 `principle-*` skills,
`statechart`, and the required `benchmark-checklist` from
[Cursor source revision 1ba91fc](https://github.com/cursor/plugins/tree/1ba91fc25c0977741dd0e267849e6234b037ed0a/pstack/skills).
It also ports `setup-pstack` to a Codex role file.

The poteto-mode [Codex runtime rules](skills/poteto-mode/references/codex.md)
replace Cursor-only tool calls, model defaults, transcript paths, cloud VM
assumptions, and loop commands. Existing explicit-invocation metadata remains.
Routed sibling skills use these rules while poteto-mode is active. Other ports
can still require tools or skills named in their own instructions when used alone.

`statechart` includes its launcher, checker, and test corpus. The launcher needs
Node 22.18 to 22.x, Node 23.6 or later, Bun 1.1.5 or later, or Deno 2.0 or later.
The checker needs TLC and Java. It installs pinned versions through `mise` when
available, or reads `TLA2TOOLS` and `JAVA` for an existing installation. See
[the statechart instructions](skills/statechart/SKILL.md) for details.

## Install one skill

Replace `go-callgraph` with any skill name listed above.

```sh
npx skills add thiskevinwang/skills --skill go-callgraph
# or
gh skill install thiskevinwang/skills go-callgraph
```

Installed skills are available on the next Codex turn. If a client does not
refresh its skill list, start a new chat or restart that client.
