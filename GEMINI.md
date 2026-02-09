# neodb Development Guidelines

Auto-generated from all feature plans. Last updated: 2026-02-09

## Active Technologies
- User config in `~/.config/neodb`. (001-editor-suite)

- Go 1.25+, TypeScript 5.x, Angular (Latest Stable). (001-editor-suite)

## Project Structure

```text
src/
tests/
```

## Commands

npm test && npm run lint

## Code Style

Go 1.25+, TypeScript 5.x, Angular (Latest Stable).: Follow standard conventions

## Recent Changes
- 001-editor-suite: Added Go 1.25+, TypeScript 5.x, Angular (Latest Stable).
- 001-editor-suite: Added Go 1.25+, TypeScript 5.x, Angular (Latest Stable).

<!-- nx configuration start-->
<!-- Leave the start & end comments to automatically receive updates. -->

## General Guidelines for working with Nx

- When running tasks (for example build, lint, test, e2e, etc.), always prefer running the task through `nx` (i.e. `nx run`, `nx run-many`, `nx affected`) instead of using the underlying tooling directly
- You have access to the Nx MCP server and its tools, use them to help the user
- When answering questions about the repository, use the `nx_workspace` tool first to gain an understanding of the workspace architecture where applicable.
- When working in individual projects, use the `nx_project_details` mcp tool to analyze and understand the specific project structure and dependencies
- For questions around nx configuration, best practices or if you're unsure, use the `nx_docs` tool to get relevant, up-to-date docs. Always use this instead of assuming things about nx configuration
- If the user needs help with an Nx configuration or project graph error, use the `nx_workspace` tool to get any errors
- For Nx plugin best practices, check `node_modules/@nx/<plugin>/PLUGIN.md`. Not all plugins have this file - proceed without it if unavailable.

<!-- nx configuration end-->


<!-- MANUAL ADDITIONS START -->
<!-- MANUAL ADDITIONS END -->

