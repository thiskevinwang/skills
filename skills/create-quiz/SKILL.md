---
name: create-quiz
description: "Create a university-level quiz document and answer key from a supplied document or RFC. Test fundamentals with multiple-choice and freeform questions, ordered subtopics, and source-line citations for every answer. Use for quiz preparation, not an interactive quiz session."
---

# Create Quiz

Create a complete quiz for later study or teaching. Test understanding of core ideas and their use, not obscure facts or section-number recall.

## Scope and model roles

The intended workflow uses a large model, such as GPT 6 Astra Ultra, to read the source and create the quiz. A smaller model, such as GPT 6.1 Sol High or GPT 6 Luna High, can later conduct it. These are workflow examples, not automatic model settings.

Produce the documents in this task. Do not start an interactive quiz, change models, or create another chat. Interactive teaching is outside this skill. The existing `$quiz-me` skill can handle that later; do not modify it.

## Establish the source

- Use the supplied document, RFC, URL, or file and any stated audience, scope, length, or format. If the source is missing or cannot be read, ask for it before writing source-dependent questions.
- Read the complete requested scope before selecting questions. For a long source, read it in parts and keep a concept-to-evidence map. Disclose unread or unavailable parts; do not claim full coverage of an excerpt.
- Record each source's title, identifier, revision or date, and original URL or file path. Preserve the requested version. Treat an Internet-Draft as that draft, not as a final RFC. Identify separately any supplementary source used to support an answer.
- Establish stable, one-based source line numbers. Prefer an immutable source with line permalinks. Otherwise save an unchanged plain-text copy or a faithful text extraction beside the quiz, give it a source ID, and number every physical line, including blank lines. Keep that copy fixed after citations are assigned.
- For HTML or PDF extraction, preserve headings, reading order, and PDF page boundaries. Check cited text against the original. State that the line numbers refer to the saved extraction, not native page or browser lines. If extraction changes a table, formula, or diagram needed for an answer, repair and verify it or omit that question.
- Supply the numbered source copy or an accessible immutable line link with the result. Temporary tool line IDs are not durable citations. A section or page number can supplement a line citation, but cannot replace it.

## Organize learning

Map the main concepts, required prior knowledge, and source evidence. Give most attention to purpose, definitions, roles, assumptions, mechanisms, and limits that explain how the subject works.

Where concepts depend on each other, organize sections in that order. A useful pattern is purpose and terms, then parts and relationships, then normal operation, then application and limits. Use only sections the source supports. Do not force unrelated concepts into a sequence.

Begin each section with a short learning objective and its prerequisite section IDs, if any. Make the progression visible without giving answers. If the user gives no length, choose enough questions to cover the main fundamentals without repetition; state the selected question count and coverage.

## Write the questions

- Include a meaningful mix of multiple-choice and freeform questions. Aim for a roughly even split unless the requested scope suggests another balance. Avoid a token question of either type.
- Move from understanding a principle to explaining or applying it. University-level difficulty should come from reasoning about the fundamentals, not specialist trivia, traps, or missing context.
- Give each question a stable ID, type, and one main learning target. Supply all scenario facts needed to answer. Do not require unstated outside knowledge.
- For multiple choice, default to four plausible options and one best answer. Use common conceptual errors as distractors. Avoid overlapping options, answer-length clues, and "all/none of the above." Vary the correct option position. Mark any requested multiple-select item explicitly.
- For freeform questions, ask for a focused explanation, comparison, or application. State the expected depth, such as a short paragraph. Avoid bundling unrelated targets into one question.
- Keep solutions, answer-revealing evidence excerpts, and grading notes out of the learner copy. Check that later questions do not state earlier answers unnecessarily.

## Build the answer key

Use the same question and section IDs. Make the key sufficient for another model or instructor to assess answers without repeating the full document analysis.

For each question, include:

- The correct option and rationale, or a concise freeform model answer.
- Essential points for a correct answer, acceptable equivalent wording, and criteria for partial credit where useful. Explain each multiple-choice distractor's error.
- A precise source-line citation beside every factual answer claim and required grading point. Cite the smallest range that establishes the claim, with source ID, section title or number, line range, and a verified link or supplied source-copy reference. Cite evidence for distractor explanations too; absence from the source does not by itself prove a claim false.
- A short hint that does not reveal the answer, plus a verified source reading pointer. Keep both in the instructor key for later interactive use.

Example citation form: `S1, section 2.1, lines 45–49 (sources/S1-numbered.txt)`. Replace the example with verified values. Prefer a verified GitHub commit permalink with a line anchor when available; never invent an anchor.

Preserve the source's conditions and strength of claims, including MUST, SHOULD, and MAY. Do not turn recommendations into requirements. For a derived answer, label it as a deduction and cite the lines that state each premise; identify supplied scenario facts separately. Do not claim that the source states the deduction directly. If the source cannot establish a required answer, revise or remove the question.

## Deliver and verify

Honor the requested document format and destination. Otherwise use a readable quiz document and a separate instructor answer key in the workspace's normal deliverable location. Include the source register and stable line references. Keep the two documents separately identifiable in any requested container or format.

Before delivery, check each answer against its cited lines, including qualifiers and exceptions. Check that each question has one matching key entry, that each multiple-choice item has the stated number of correct options, and that all required grading points have evidence. Check the mix of types, prerequisite order, and coverage of fundamentals.

Add a brief handoff note in the instructor key: a later facilitator can use `$quiz-me` with the quiz, key, and sources. Its current open-question policy can require adapting a multiple-choice item into an open question with the same learning target. The document's mixed question types remain intact. Do not implement the interactive session here.

Return links to the quiz, key, and source material, with any source-access or coverage limits. Do not paste the answer key into the completion message unless requested.
