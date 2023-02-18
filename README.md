# Form3 Homework

Name: Adan Jesus Suarez Garcia

Form3 homework.

Ref:

https://github.com/form3tech-oss/interview-accountapi

## Client library

The library is a Go library to be use as so. For simplicity, it implements only limited functionality of the Account endpoints of Form3.

For accounts, it implements:
- `Create` a new bank account.
- `Fetch` a bank account.
- `Delete` a bank account.

## Tests
The description says the following: `Be well tested to the level you would expect in a commercial environment. Note that tests are expected to
run against the provided fake account API.`

To fullfil this requirement, I implemented two groups, unit tests and integration tests. The integration test are the one that will run against
the provided fake account API.
- The unit tests are located in each module.
- The integration tests are in a specific folder called `integration`.

My tests in the integration folder include "contract" checking. It could be seen as a combination of integration and contract testing.

**Important note:** During the developing I found some discrepancies between the Documentation and the behavior of the fake API, therefore some tests fail. All integration tests are documented with the expected behavior.

## Mocks

I created my unit tests using mocks generated by [Mockery](https://github.com/vektra/mockery). There are two opposite approach related to mocks, there are teams that commit them to the repository and there are team that don't. There are pros and cons, so it is usually a matter of team decision. The ones in favour say that when committing them you eliminate the needs for generate them, that sometimes require external libraries/software. The others, consider that mocks are a tool to help with the tests, they should not be in the repo, so you generate them when you needed them.

I this case I didn't commit the mocks, but the unit tests doesn't pass without them. In the `Dockerfile` I included a `RUN` conditioned to success on the installation of Mockery, in a way that, if it fails, the unit tests will not run to avoid the failure of running all tests: unit and integration. The requirement of the exercise is to see the tests running against the fake API (integration test).

If the Mockery installation fails for any reason when creating the image after `docker-compose up`, you still can to run the unit tests at `docker-compose up`:
- Get Mockery. There are different instructions depend on your system/preference [link](https://vektra.github.io/mockery/installation/#github-release)
- Generate the mocks based on the instruction. You would need to add the following Mockery flag: `--inpackage` to generate the mocks, only if the way you chose wasn't `go generate ./...`
- Uncomment line 16 in Dockerfile. It has a note.
- Run `docker-compose up`.

## Vendor

Like with mocks, there is a discussion of committing the dependencies or not to your repo. At the end, it depends of the team.
I saw in your client [repo](https://github.com/form3tech-oss/go-form3) you included it, therefore I took the same approach.

# Production readiness
## Configuration

This implementation of client library needs basically to parameters to run. The `Form3 URL` and the `account path`.
To set those parameters I implemented two different ways. One is setting them as parameters when instantiate form3 object and the other is read them from the environment variables. For production ready, it should include a third case that could be read them from a file like `yaml` or `toml`. I decided to not implement that case because it will require either implement a parser for those files or include a third-party library that I am not allowed based on the instruction.

## Retry mechanism

The API documentation encourage us to implement a [retry mechanism](https://www.api-docs.form3.tech/api/schemes/sepa-instant-credit-transfer/introduction/timeouts/retry-strategy) on failure. This could be implemented in the client library. I did some test running 10000 concurrent requests against the fake API and instead of receiving 429 I received 500 and, in my case, the dockers got unresponsive after that so I couldn't test it properly and all the following tests never passed, so I decide to not implement it for simplicity, but it is definitely a needed feature for production readiness.

## Headers

The exercise requirements say to not implement authentication so I didn't do anything about it. For production ready we should include the Authentication header.

## Create

The instruction require to be simple and concise, so only implemented tha basic functionality for `Create`. One nice feature could be the possibility to pass parameters for creation instead of passing the `model.Data` The values that the user doesn't change often, could be set through configuration make it even simpler.
