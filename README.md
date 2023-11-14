# ae-cli-tool

## Description
AE-CLI-Tool is a versatile command-line interface tool specifically designed for Analytics Engineers. It aims to automate and streamline various CLI commands essential in data analytics workflows, with a key feature being the robust automation of `kubectl port-forward` commands. This tool enhances productivity and stability in Kubernetes environments, making it an indispensable utility for Analytics Engineers.

## Features
- **Kubernetes Port Forwarding**: Automates and maintains stable `kubectl port-forward` commands.
- **Custom CLI Command Automation**: Simplifies repetitive CLI tasks specific to data analytics.
- **Cross-Platform Compatibility**: Works seamlessly on different operating systems.

## Installation

### Prerequisites
- Go 1.15 or higher
- Access to a Kubernetes cluster (for Kubernetes-related features)

### Steps
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/larryfisherman25/ae-cli-tool.git
   cd ae-cli-tool
   ```

2. **Build the Binary**:

   ```bash

    go build
   ```

## Usage
### Default Kubernetes Port Forwarding

To use the default port forwarding settings:

  ```bash

  ./ae-cli-tool forward -d
  ```

### Custom Port Forwarding

For custom port forwarding settings, provide the service and port mapping:

  ```bash

./ae-cli-tool forward [service] [localPort:remotePort]
  ```

### Contributing

We welcome contributions to the AE-CLI-Tool! If you have suggestions for improvements or bug reports, please open an issue. If you'd like to contribute code, please open a pull request with a clear description of your changes.
Versioning

We use Semantic Versioning for our tool. For the available versions, see the tags on this repository.
License

This project is licensed under the MIT License - see the LICENSE file for details.
Acknowledgments

    Special thanks to everyone who contributed to the development and maintenance of this tool.


### Notes:
- **Customization**: You may need to adjust some sections to better fit your project's specifics, especially under "Features" and "Additional Commands".
- **License Link**: Replace `[LICENSE](LICENSE)` with the actual link to your license file in the repository.
- **Repository URL**: Make sure to replace `https://github.com/larryfisherman25/ae-cli-tool.git` with the actual URL of your repository.
- **Acknowledgments**: You can add or remove acknowledgments as per your preference.
