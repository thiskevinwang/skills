### Session pickup

**You own the resume point. Read the prior trail, don't redo it.**

1. Locate the prior trail from known Codex chat IDs, task attachments, a transcript path provided for this task, or a pushed branch. Use Codex chat-read tools for recent summaries, then read older turns only when needed. Do not search unrelated chats or global session directories. Use a subagent for a long transcript when available, with the complete scope and model policy from the Subagents section.
2. Reconstruct operational state. The branch and worktree, what already landed (`git log`, `git diff` against the base), the open todos, the decisions made. Use the prior trail as historical evidence. Check changeable claims against the current artifact when needed. Resist the bias to re-derive it.
3. Diff done vs pending. Compare what shipped against what was planned, name the resume point, do not re-run the prior repro or redo completed work. Repeat a prior check only when changed state or missing evidence makes it necessary.
4. Route the remaining work to the matching playbook and pick the verdict: continue the execution, ship a finished recommendation, ratify or override a prior conclusion, or postmortem a failed run. The pickup playbook ends here. The routed playbook owns the rest.
5. Verify the inherited claims against the original goal on the real artifact (the **principle-prove-it-works** skill). A passing prior self-report is not the proof.

**Reply:** where the prior agent stopped, what you inherited vs redid (ideally nothing redone), the resume point, and the outcome.
