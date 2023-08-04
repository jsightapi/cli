<div align="center">

<div>  
  &nbsp; 
</div>
	
<a href="https://jsight.io" align="left"><img src="./img/jsight-logo.svg" alt="JSight" width="148px"/></a>

# JSight CLI

  [![Golang](https://badges.aleen42.com/src/golang.svg)](https://go.dev/)
  [![Telegram support](https://img.shields.io/badge/Support-Telegram-blue)](https://t.me/jsight_support)
  [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](./CONTRIBUTING.md)
  [![License](https://img.shields.io/github/license/jsightapi/jsight-server?colorB=ff0000)](./LICENSE)
  [![JSight on Facebook](https://img.shields.io/badge/Facebook-1877F2?logo=facebook&logoColor=white)](https://www.facebook.com/jsightapi)
  [![JSight on LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?logo=linkedin&logoColor=white)](https://www.linkedin.com/company/jsightapi/)
  [![Twitter Follow](https://img.shields.io/twitter/follow/jsightapi.svg?style=social)](https://twitter.com/jsightapi)

  <a href="https://www.alchemistaccelerator.com/portfolio?class=29"><img width="300px" src="./img/alchemist.svg" alt="Alchemist Accelerator"/></a>

<div>
  &nbsp;
</div>

  :star: **Star us on GitHub — it motivates us a lot!** :star:

<div>
  &nbsp;
</div>
</div>

**JSight CLI** is the cli tool which performs various task when working with [JSight API
language](https://jsight.io/docs/jsight-api-0-3).

Currently, the **JSight CLI** allows you to perform only one task:

1. Parsing code in the JSight API language and converting it to HTML document file.

> :fire: If you have any ideas or suggestions, please share with us:
> 
> - Email: [support@jsight.io](mailto:support@jsight.io)
> - Telegram: [@jsight_support](https://t.me/jsight_support)

## :rocket: &nbsp; Install

Download `jsight` cli tool from the official website and move to the `/usr/local/bin` folder:

```sh
wget https://jsight.io/downloads/jsight-cli/1.0.0/linux-x64/jsight
mv jsight /usr/local/bin
chmod +x /usr/local/bin/jsight
```

Check the installation:

```sh
jsight version
```

## :rocket: &nbsp; Usage

- `jsight version` — outputs the current version of JSight CLI.
- `jsight help` — outputs the manaul fo JSight CLI.
- `jsight doc html <jsight file>` — parses the `<jsight file>` and outputs the corresponding html
  document (or an error description).

Example:

```sh
jsight doc html my-api-spec.jst > my-api-spec.html
```

## :rocket: &nbsp; Manual Building and Testing

- `make` — builds binary to the `./build` folder and runs the tests.
- `make build` — builds binary to the `./build` folder.
- `make test` — runs the tests against the `./build/jsight` binary.

## :bookmark_tabs: &nbsp; Versioning

The **JSight CLI** adheres to the [Semantic Versioning 2.0.0](https://semver.org/).

```
{MAJOR version}.{MINOR version}.{PATCH version}
```

The JSight CLI version is specified in the file
[main.go](./main.go) in the global variable `Version`.

## :notebook_with_decorative_cover: &nbsp; Dependencies

All the dependencies are described in the file [go.mod](./go.mod).

<div>
  &nbsp;
</div>

## :sunglasses: &nbsp; Contributing

Contributing is more than just coding. You can help the project in many ways, and we will be very
happy to accept your contribution to our project.

Details of how you can help the project are described in the [CONTRIBUTING.md](./CONTRIBUTING.md)
document.

<div>
  &nbsp;
</div>

### Contributors

<a href="https://github.com/dshemin"><img src="https://avatars.githubusercontent.com/u/11780307?v=4" width="100" height="100" alt=""/></a>
<a href="https://github.com/Emptyfruit"><img src="https://avatars.githubusercontent.com/u/14968783?v=4" width="100" height="100" alt=""/></a>
<a href="https://github.com/add2"><img src="https://avatars.githubusercontent.com/u/3954234?v=4" width="100" height="100" alt=""/></a>
<a href="https://github.com/constantine-malyshev"><img src="https://avatars.githubusercontent.com/u/101567029?v=4" width="100" height="100" alt=""/></a>

<div>  
  &nbsp; 
</div>

## :speech_balloon: &nbsp; Bugs and Feature Requests

Do you have a bug report or a feature request? 

Please feel free to add a [new
issue](https://github.com/jsightapi/cli/issues/new) or write to us in support:

- Email: [support@jsight.io](mailto:support@jsight.io)
- Telegram: [@jsight_support](https://t.me/jsight_support)

<div>  
  &nbsp; 
</div>

## :grey_question: &nbsp; Support

If something is unclear to you, please contact support; we try to respond within 24 hours. Moreover,
it is critical for us to understand what is unclear from the first instance.

- Email: [support@jsight.io](mailto:support@jsight.io)
- Telegram: [@jsight_support](https://t.me/jsight_support)

<div>  
  &nbsp; 
</div>

## :receipt: &nbsp; License

This project is licensed under the Apache 2.0 License. See the [LICENSE](./LICENSE) file for more
details.

<div>  
  &nbsp; 
</div>

## :book: &nbsp; Resources

- JSight Official Website: https://jsight.io.

### Documentation

- JSight API Language Quick Tutorial: https://jsight.io/docs/jsight-api-0-3-quick-tutorial.
- JSight API language specification on the official website: https://jsight.io/docs/jsight-api-0-3.  
- JSight API language specification on GitHub:
  https://github.com/jsightapi/specification/tree/main/versions/JSight%20API.

### Publications

- JSight blog: https://jsight.io/blog.
- Official Facebook page: https://www.facebook.com/jsightapi.
- Official Twitter: https://twitter.com/jsightapi.
- Official Linkedin: https://www.linkedin.com/company/jsightapi.

### Others

- All JSight repositories: https://github.com/jsightapi.

<div>  
  &nbsp; 
</div>

## :handshake: &nbsp; Partners

- We have successfully completed [class #29 of the Alchemist
  Accelerator](https://www.alchemistaccelerator.com/portfolio?class=29).

<div>  
  &nbsp; 
</div>

## :trophy: &nbsp; Acknowledgments

We sincerely thank all those without whom this project would not have been possible:

- [Alchemist Accelerator](https://www.alchemistaccelerator.com/),
- [Urfave](https://github.com/urfave) for his [cli](https://github.com/urfave/cli) library,
- [Lucas Jones](https://github.com/lucasjones) for his
  [reggen](https://github.com/lucasjones/reggen) library.


<div align="center">

:star: **Star us on GitHub — it motivates us a lot!** :star:

</div>
