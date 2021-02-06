# CHANGELOG

## v2.0.0 (2021-02-05)

* Overhauled properties of go objects (eg: `Id` to `ID`, `Url` to `URL` etc) - breaking change
* Added GitHub Actions
* Bumped Go version from 1.14 to 1.15
* Added missing user-agent to headers (closes #2)
* Linted the entire project and used gofmt for better formatting

## v1.0.0 (2020/04/17)

* Removed lots of duplicate code
* Use maps properly and return data instead of assuming the user wanted to print instead
* Use built in structures
* Better examples and documentation

## v0.2.0 (2020/03/25)

* Added Inventory actions
* Added Location actions
* Added Ticket actions
* Consolidated and re-used code in a proper client/response struct and function
* Added a complete set of examples
* This release is considered the MVP

## v0.1.0 (2020/03/24)

* Working customer actions
* Added documentation

## v0.0.1 (2020/03/23)

* Initial release, can retrieve customer records