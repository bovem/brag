# brag
Inspired by Julia Evan's blog [Get your work recognized: write a brag document](https://jvns.ca/blog/brag-documents/)
I created this tool to keep a record of my brags in Markdown Format and 
refer them later to create a formatted bragging document.

A bragging or hype document is created to keep a record of achievements over time.
This document/information extracted from this document could be shared with managers 
or used for self-reflection and improvement. I refer to it during my daily standups
and weekly sync-up meetings.

Before this, I was using Google Docs to maintain a record of my accomplishments
but I decided to create a command line tool because:
* I don't have to leave my development environment to add a new accomplishment. 
* I edit and maintain the document using applications like VSCode, VIM, and Obsidian.
* I can use Git for Version Control. 

## Setup and Usage
To initialize a bragging document directory create a Git repository on any provider
like GitHub, GitLab, BitBucket, etc., and clone the repository locally. You can also
create a local Git repository using `git init`.

The location of the bragging document directory is specified using the `BRAG_DOCS_LOC`
environment variable.
```bash
export BRAG_DOCS_LOC=/work-docs
```

Once the directory is available locally you can initialize the document directory
using the `init` subcommand.
```bash
brag init
```

You can add your brags using the `-c`/`--comment` flag.
```bash
brag -c "YOUR TEXT HERE"
```

The brags will be added to the directory specified in the `BRAG_DOCS_LOC` environment
variable.

To review brags added in a given time period use subcommand `about`.
```bash
brag about last-week
```

If the `BRAG_DOCS_REPO_SYNC` environment variable is set to `true` then new brags
will be committed and pushed to the Git remote.

## Todos
- [ ] Add a `edit` subcommand to edit recent entries
- [ ] Add a `summarize` subcommand to create a formatted bragging document using LLMs hosted on Ollama

## Sample Output
```text
2024-01-02                                                                                                                                        
* 11:12:50      Updated Quarterly Status Report
* 17:17:42      Knowledge transfer presentation with new interns

2024-01-03
* 12:08:36      Filed an issue with UI text rendering on the pre-production environment. ID:12345

2024-01-04
* 17:39:50      Provisioned a new Kubernetes cluster
```
