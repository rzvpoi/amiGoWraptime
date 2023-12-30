# amiGoWraptime
<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/rzvpoi/amiGoWraptime">
    <img src="images/logo.jpg" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">amiGoWraptime</h3>

  <p align="center">
    Asterisk Wrap-up Time manager


   <a href="https://github.com/rzvpoi/amiGoWraptime/issues">Report Bug</a>
    ·
    <a href="https://github.com/rzvpoi/amiGoWraptime/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
      </ul>
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
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

This application is designed to address the issue of Asterisk Wrap-up Time by automatically initiating a pause status for agents in a queue once they have completed a call for a set amount of time. 

The purpose of this feature is to optimize post-call activities and streamline workflow efficiency, ensuring that agents have the necessary time and resources for wrap-up tasks without manual intervention. 

By implementing this solution, the application aims to enhance overall call center performance and improve the handling of post-call processes within the Asterisk system.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Built with Go

The application was built with Go and uses the [amigo](https://pkg.go.dev/github.com/ivahaev/amigo) package to connect to Asterisk AMI

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

AmiGoWraptime can be run either in a docker container or locally on a machine. 
<br><br>
Before starting with the installation process make sure that the server where the app will be installed can connect to the AMI instance.

### Installation

#### Run it with Docker
1. Download the Docker image
```bash
docker pull trp8818/goamiwraptime:latest
```
2. Run the image. Change the env variables for you case
```bash
docker run -d \
  --name goamiwraptime \
  -e AMI_HOST="localhost" \
  -e AMI_USERNAME="admin" \
  -e AMI_PASSWORD="password" \
  -e QUEUES="300,500" \
  trp8818/goamiwraptime:latest
```
#### Run it locally on a machine
1. You need to download the right executable from the releases page
The latest version right now is [v1.0.0](https://github.com/rzvpoi/amiGoWraptime/releases/tag/v1.0.0)
For linux and macos you can run the following command:
```bash
curl -O https://github.com/rzvpoi/amiGoWraptime/releases/download/v1.0.0/gowraptime_linux_v1.0.0_amd64
``` 
For windows powershell:
```powershell
Invoke-WebRequest -Uri "https://github.com/rzvpoi/amiGoWraptime/releases/download/v1.0.0/gowraptime_windows_v1.0.0_x64.exe" -OutFile "gowraptime_windows_v1.0.0_x64.exe"
```

2. After the download finishes you need to create an .env file with the following text
```vim
AMI_HOST = 'localhost'
AMI_USERNAME = 'admin'
AMI_PASSWORD = 'password'
QUEUES = "300,500" 
WRAPTIME = "30"
```
* QUEUES: You can write one or more queues depending from which queue you want agents to be put in pause
* WRAPTIME: this variable is in seconds and the default value is 30 seconds.

3. After this you can start the app and watch it work.

<!-- USAGE EXAMPLES -->
## Usage
This application is designed to automatically pause agents in the queue for a predefined wrap-time, eliminating the need for manual intervention.
<br>
It is particularly effective when integrated with an autodialer or predictive dialer system.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Poienariu Razvan - [LinkedIn](https://ro.linkedin.com/in/razvan-poienariu) - razvanpoienariu@gmail.com

Project Link: [https://github.com/rzvpoi/amiGoWraptime](https://github.com/rzvpoi/amiGoWraptimee)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


[contributors-shield]: https://img.shields.io/github/contributors/rzvpoi/amiGoWraptime.svg?style=for-the-badge
[contributors-url]: https://github.com/rzvpoi/amiGoWraptime/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/rzvpoi/amiGoWraptime.svg?style=for-the-badge
[forks-url]: https://github.com/rzvpoi/amiGoWraptime/network/members
[stars-shield]: https://img.shields.io/github/stars/rzvpoi/amiGoWraptime.svg?style=for-the-badge
[stars-url]: https://github.com/rzvpoi/amiGoWraptime/stargazers
[issues-shield]: https://img.shields.io/github/issues/rzvpoi/amiGoWraptime.svg?style=for-the-badge
[issues-url]: https://github.com/rzvpoi/amiGoWraptime/issues
[license-shield]: https://img.shields.io/github/license/rzvpoi/amiGoWraptime.svg?style=for-the-badge
[license-url]: https://github.com/rzvpoi/amiGoWraptime/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://ro.linkedin.com/in/razvan-poienariu