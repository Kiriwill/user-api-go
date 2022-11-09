# Contributing guide
For whom it may concern this is the contributing guide to the basemed team, and in any doubt feel free to contact us through *slack/email/teams*.
*This is a template file, and can differ a bit deppending on the project.*

Consider it a guide to newcomers. *A guide with our project patterns and considerations*. 

## Getting started
### Project structure
- `/cmd` will have all the code that will be built and shipped as binaries. they are most likely to be named `main.go` and be found at `/cmd/<executable-name>/main.go`
- `/pkg` will have the actual shared codebase of the project. all of the modules used by the entrypoints will come from this folder
- `/scripts` will have useful utilitary scripts such as `dev.sh` that usually runs the project at a dev-local environment
- `/pkg/*<storage|repository>*/migrations` (optional) this one is optional and if found it will contain eventual migrations. We use the [migrate CLI](https://github.com/golang-migrate/migrate) to create and manage our migrations
- `/misc` (optional) will have miscellaneous scripts and files that do not relate to the project directly but are (or were) used to build and maintain the project. (such as a migration script e.g.)

### Versioning
This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
To version a change you should use the concepts provided by the link above.
You can tag it from the following GIT CLI command `git tag v1.2.3 -m "release version 1.2.3"`.

## Style Guides

### Pull Requests
#### Creating a pull request
The pull request must have as the title the following information: **(PR TYPE): (ISSUE CODE) | (Short description)**

Then you should fill the default Pull Request Template with relevant information about the changes you are doing. Some fields are not mandatory and are marked as so.

The last part of the pull request template is a checklist that will be enforced through a github action to check if you *agreed to all of the terms*.

#### Is your pull request not final?
In case of a not-final pull request you can merge it into a feature branch.
With this you will be grouping all the code needed for a specific functionality, that can be tested as a whole at the staging environment.
**As we do not use gitflow, do not expect to find a development branch.**

#### Review a pull request
A **ready to review** pull request is open for review not only of those marked as review required. So feel free to contribute to the project and review the changes.
Always remember to make your comments grouped by a review.
*Being respectfull is a basic requirement*.

#### Merge a pull request
Congratulations! if your changes are about to be merged to the **main** branch expect it to be available as soon as possible through our continuous delivery workflows. (***subject to change soon!!***)

### Commits
#### Commit message
It is important to keep things consistent. We're based off of english-only programming languages, which will inevitably have much clearer and better documentation in english, with all it's libraries, commit messages, issues and such in english. So keep the commit messages in english.
The commit message must be in english, with a simple description of what is included on the commit.

Example: `git commit -m "fixed the bug where the thing didn't do it's thing on that other thing"`

#### Commit structure
Try to structure your commits by feature, changeset or context.
If you changed a LOT of things, try to do small commits, which you can achieve with `git add --patch --interactive` where you'll add to the HEAD only some parts of the files, and you can control much more what you'll add.
