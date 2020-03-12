# Releases

## Versioning nomenclature

Version numbers for releases of this project are intended to be consistent with the [Semantic Versioning 2.0.0](https://semver.org/) specification. In short:

> Given a version number MAJOR.MINOR.PATCH, increment the:
>
> 1. MAJOR version when you make incompatible API changes,
> 1. MINOR version when you add functionality in a backwards compatible manner, and
> 1. PATCH version when you make backwards compatible bug fixes.
>
> Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.

## Tagging the release

Tagging process:

```
git checkout develop
git checkout -b release-xxx (e.g., release-120)
git tag vx.x.x (e.g., v1.2.0)
git push origin vx.x.x
git push origin release-xxx
```

Once tagged, go to GitHub:

- Click `New pull request`
- Change base to `master`, set compare to `release-xxx`
- Create pull request
- Review commits you may not have had an opportunity to review during the sprint.
- Wait for tests to finish running
- Merge the pull request, delete the `release-xxx` branch

## Creating the release

- Go to Releases tab from main page.
- Create new release, type in the tag name – v1.2.0. It will appear in drop down. (i.e., you won't be creating it in this UI, you wont specify any particular branch)
- Type in release notes. You can get a template by going to previous release and click Edit, then copy the text there. Keep the highlights high-level, consolidating similar tickets into a basic description.

## Closing out the sprint project (board)

At the end of the sprint, on the sprint project page:

- Ensure that all cards in `To Do` are moved to next sprint or removed from the project.
- Ensure that all cards In Progress are moved to next sprint, then moved to In Progress from To Do
- Close sprint project (via hamburger menu on far right)

## Announce the release

Once you publish the release, post an announcement to the corresponding Microsoft Teams channel, tagging the appropriate stakeholders. The announcement should have a link to the release on GitHub, and the contents of the release notes.

## Create a new sprint project (board)

- One the Project page, click `New Project`.
- Use the date of the last day of the sprint as the project name.
- Use the `Automated kanban` template.
- Click `Create project`.
- On the new sprint project, delete any default cards in the `To do` column.
- Click the ellipsis in the `Done` column and uncheck the automation options for pull requests, leaving just the automation for issues.
