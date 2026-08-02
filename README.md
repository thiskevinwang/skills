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

### `graph-this-out`

Maps moderate- to high-complexity implementation work as a dependency-aware
Mermaid DAG, including the critical path and safe opportunities for parallel
agents.

Install it with:

```sh
npx skills add thiskevinwang/skills --skill graph-this-out
# or
gh skill install thiskevinwang/skills graph-this-out
```

After installation, start a new Codex task or restart Codex if the skill does
not appear immediately.
