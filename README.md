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

- [graph-this-out](skills/graph-this-out/SKILL.md): Map work as a dependency graph.
- [go-callgraph](skills/go-callgraph/SKILL.md): Build and display Go call graphs.
- [create-quiz](skills/create-quiz/SKILL.md): Create a quiz and answer key with source citations.
- [quiz-me](skills/quiz-me/SKILL.md): Learn concepts with one question at a time.

### Principles

These 21 skills preserve the installed instructions and metadata. In Codex,
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

### Cursor plugin ports

These installed ports come from [pstack](https://github.com/cursor/plugins/tree/main/pstack),
except [thermo-nuclear-code-quality-review](skills/thermo-nuclear-code-quality-review/SKILL.md),
which comes from [thermos](https://github.com/cursor/plugins/tree/main/thermos).

- [architect](skills/architect/SKILL.md)
- [arena](skills/arena/SKILL.md)
- [automate-me](skills/automate-me/SKILL.md)
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

Some workflows still refer to Cursor tools, transcripts, or other skills.
Those dependencies must be available to use those workflows.

## Install one skill

Replace `graph-this-out` with any skill name listed above.

```sh
npx skills add thiskevinwang/skills --skill graph-this-out
# or
gh skill install thiskevinwang/skills graph-this-out
```

After installation, start a new Codex task or restart Codex if the skill does
not appear immediately.
