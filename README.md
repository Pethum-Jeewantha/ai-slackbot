# AI Slackbot

[![GitHub][linkedin-shield]][linkedin-url]
[![GitHub][contributors-shield]][contributors-url]
[![GitHub issues][issues-shield]][issues-url]
![GitHub][license-shield]
![GitHub Repo stars][stars-url]
![GitHub forks][forks-url]


<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

## About The Project

This is a Slack chatbot built with Golang and integrated with AI capabilities. The chatbot uses AI to understand and respond to user messages in a conversational manner. The code includes a Golang implementation of Slack's API for receiving and sending messages, as well as libraries for integrating with third-party AI services like [wit-go](https://github.com/wit-ai/wit-go), [wolfram](https://www.wolframalpha.com). Additionally, the repository includes documentation to help developers get started with building and deploying their own AI-powered Slack chatbots.

## Getting Started

### Installation

1. First clone the repository.

   `https://github.com/Pethum-Jeewantha/ai-slackbot.git`

2. Once cloned, open the repository from an IDE

4. Make a `.env` file in the project root. Next, include the attributes below with values. Pre-configuring these tokens from relevant sources is necessary.
   - SLACK_APP_TOKEN
   - SLACK_BOT_TOKEN
   - WIT_AI_TOKEN
   - WOLFRAM_APP_ID

5. Install any necessary dependencies using `go get` or another package manager.

6. Run the project using `go run ./main.go` after the building is complete.

## Usage

For programmers looking to utilize Golang to create a Slack chatbot with AI capabilities, this chatbot offers a ready-to-use option. The chatbot can comprehend user messages and have natural-sounding conversations with users thanks to the integration of AI. Offering a pre-built basis for creating AI-powered chatbots that can be customized to match their particular use case, this may save developers a lot of time and effort.

## License

Copyright &copy; 2023 - present Pethum Jeewantha. All rights reserved.

Licensed under the [MIT](LICENSE.txt) license.

## Contact

Pethum Jeewantha - [Twitter](https://twitter.com/JeewanthaPethum?s=08) - me@pethumjeewantha.com

## Acknowledgements

* [Slacker](https://github.com/shomali11/slacker)
* [Wit-go](https://github.com/wit-ai/wit-go)
* [Wolfram](https://www.wolframalpha.com)
* [Gjson](https://github.com/tidwall/gjson)
* [Img Shields](https://shields.io)
* [Choose an Open Source License](https://choosealicense.com)

[contributors-shield]: https://img.shields.io/github/contributors/Pethum-Jeewantha/ai-slackbot

[contributors-url]: https://https://github.com/Pethum-Jeewantha/ai-slackbot/graphs/contributors

[issues-shield]: https://img.shields.io/github/issues/Pethum-Jeewantha/ai-slackbot

[license-shield]: https://img.shields.io/github/license/Pethum-Jeewantha/ai-slackbot

[issues-url]: https://github.com/Pethum-Jeewantha/ai-slackbot/issues

[stars-url]: https://img.shields.io/github/stars/Pethum-Jeewantha/ai-slackbot?style=social

[forks-url]: https://img.shields.io/github/forks/Pethum-Jeewantha/ai-slackbot?style=social

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat&logo=linkedin&colorB=555

[linkedin-url]: https://www.linkedin.com/in/pethum-jeewantha
