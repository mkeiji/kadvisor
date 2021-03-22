# Client

## Generate a library

Run `nx g @nrwl/react:lib my-lib` to generate a library.

> You can also use any of the plugins above to generate libraries as well.
> Libraries are sharable across libraries and applications. They can be imported from `@client/mylib`.

## Generate a component

Run `npx nx g @nrwl/react:component my-new-component --project=my-lib-or-project-app --export` to generate a new component.

> Note: `my-new-component` is the component folder/name and `my-lib-or-project-app` is the project/folder you want the component inside

## Run Development server

Run `nx serve kadvisor-app` for a dev server. Navigate to http://localhost:4200/. The app will automatically reload if you change any of the source files.

## Build

Run `nx build kadvisor-app` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `nx test kadvisor-app` to execute all the unit tests via [Jest](https://jestjs.io).
Run `nx affected:test` to execute the unit tests affected by a change.
Run `nx test kadvisor-app --testFile <<nameOfTheFile>>` to execute unit tests of a single file

## Running end-to-end tests

Run `ng e2e kadvisor-app` to execute the end-to-end tests via [Cypress](https://www.cypress.io).
Run `nx affected:e2e` to execute the end-to-end tests affected by a change.

## Understand your workspace

Run `nx dep-graph` to see a diagram of the dependencies of your projects.
