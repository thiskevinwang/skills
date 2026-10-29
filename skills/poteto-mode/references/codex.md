# Codex runtime

Read this file before a playbook. Apply these rules to routed skills too. They replace Cursor runtime assumptions in the source workflows.

## Tools and skills

- Invoke skills as `$skill-name`, or read their `SKILL.md`. Paths that start with `scripts/`, `playbooks/`, or `references/` use the owning skill directory. A path that starts with `../` inside a playbook is relative to that playbook file. Resolve sibling skills from the installed skills directory. A path such as `pstack/skills/swarm/SKILL.md` means the sibling `swarm` skill. Do not assume the current project contains pstack.
- Check the available tools and their schemas before use. Use Codex subagent tools such as `spawn_agent`, `send_message`, and the available wait tool. Cursor `Task`, `subagent_type`, `run_in_background`, and `environment: "cloud"` are not Codex tool arguments.
- Use `worker` for edits and `explorer` for code questions when those agent types exist. No `poteto-agent` installation is required. Give each delegate the relevant skill paths, goal, write scope, constraints, and required evidence. State that other workers can edit the same checkout and that their edits must be retained.
- Run independent delegates within the session's actual limit. Queue the remaining work. A cloud VM is not implied by a subagent. Use separate worktrees and ports when workers need separate state. If subagents are unavailable, run the same checks in sequence and state that independent review remains pending.
- Use an available question tool only for missing facts, preferences, or required approval. Use the current plan tool for task steps. If absent, keep a short checklist in the chat or a task file.
- For browser or desktop proof, use the available browser or computer tools. For a CLI, use the terminal or a pseudo-terminal. A `control-ui` or `control-cli` skill is optional. If no driver exists, name the missing proof and use the nearest runnable check.
- For code cleanup, use `deslop` if installed. Otherwise inspect the diff for unused code, redundant guards, duplicate logic, and needless abstractions. Use `$skill-creator` for skill edits when available.

## Models

Read `${CODEX_HOME:-$HOME/.codex}/pstack-models.md` if it exists. `$setup-pstack` maintains this optional role file. Use only model IDs and reasoning levels accepted by the current session. Never translate Cursor model names or speed suffixes into guessed Codex IDs.

A `reasoning effort for <role>` line can set a separately confirmed level. Without a valid override, inherit the parent model and reasoning effort. Omit model arguments. `auto` and `inherit-parent` mean omission, not literal model IDs. Where the tools require a fresh context for an override, select it and include the complete brief. Do not alter the service tier or global configuration to imitate a Cursor speed choice.

Different prompts on one model can provide separate reviews. They do not provide evidence of agreement across model families. If a requested model is rejected, report the limit, use parent inheritance, and do not repeat the rejected ID.

## Wakeups and records

Use the available agent or process wait tools for active work. Keep waits short enough to report useful progress. A watcher running in a terminal cannot wake a chat after its turn ends.

For a requested reminder, monitor, or recurring run, use the Codex automation tool with a heartbeat attached to this chat. Check for an existing schedule before creation. Keep quiet when the state has no meaningful change. Do not create an automation merely because a playbook mentions a loop. If scheduling is unavailable, state that limit and save a resume point. Do not claim a future check is active.

Keep program state in a task-local `notes.md` or an existing project tracker. Record owner IDs, write scopes, branches, head revisions, dependencies, and evidence paths. Codex does not require Cursor's Project store. Share these pointers through subagent tools. Do not create user-owned chats for internal delegates.

For session pickup and evaluations, use known chat IDs with Codex chat-read tools, or a transcript path provided for this task. Read recent summaries first. Do not search unrelated chats or global session directories. If tool traces are unavailable, mark chain-following as unverified.

## Scope and delivery

The user's request and current permissions control all actions. A playbook does not grant permission to send messages, publish, merge, deploy, remove data, or change unrelated skills. Use existing authorization. Do not request the same approval twice.

Use the actual PR base or repository default branch. Never assume `main`. Reuse a suitable worktree. Preserve unrelated edits. Never use a reset to replace a worker's checkout. For a new worktree, use the Codex worktree tools when available.

Create a PR when the task calls for delivery through a PR. Follow repository instructions for its body and draft state. After creation, attach every PR with `attach_artifact` when available. A local skill edit alone does not require a push or PR.

Read current attachments and chat state before worktree cleanup. Archive managed worktrees with the Codex archive tool. A clean Git status does not prove that ignored files or active processes can be removed. Hold a candidate when its use or recoverable contents are unknown.

## Sources

- [Codex subagents](https://learn.chatgpt.com/docs/agent-configuration/subagents).
- [Cursor pstack source](https://github.com/cursor/plugins/tree/1ba91fc25c0977741dd0e267849e6234b037ed0a/pstack).
