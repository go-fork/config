# Security Policy

## Supported Versions

We actively support the following versions of go.fork.vn/config with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |
| < 1.0.0 | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security vulnerability in go-fork/config, please follow these steps:

### Private Disclosure

1. **Do not create a public GitHub issue** for security vulnerabilities
2. Send an email to security@fork.vn with:
   - A clear description of the vulnerability
   - Steps to reproduce the issue
   - Potential impact assessment
   - Any suggested fixes (if available)

### What to Include

Please include the following information in your report:
- **Description**: Clear description of the vulnerability
- **Impact**: Potential security impact and affected components
- **Reproduction**: Step-by-step instructions to reproduce
- **Environment**: Go version, OS, package version
- **Additional Context**: Any additional information that might be helpful

### Response Timeline

- **Initial Response**: Within 48 hours of receiving the report
- **Status Update**: Weekly updates on investigation progress
- **Resolution**: Target resolution within 30 days for high-severity issues

### Security Update Process

1. We will investigate and validate the reported vulnerability
2. We will develop and test a fix in a private repository
3. We will coordinate with the reporter on disclosure timeline
4. We will release a security update with appropriate version bump
5. We will publish a security advisory with details

### Scope

This security policy covers:
- The main go-fork/config package
- Configuration handling and validation
- Integration with go-fork/di dependency injection
- File system operations related to configuration

### Out of Scope

The following are typically out of scope:
- Vulnerabilities in dependencies (report to respective maintainers)
- Issues requiring physical access to the system
- Social engineering attacks
- DoS attacks that require excessive resources

## Security Best Practices

When using go-fork/config in your applications:

1. **File Permissions**: Ensure configuration files have appropriate permissions
2. **Sensitive Data**: Never store passwords or API keys in configuration files
3. **Environment Variables**: Use environment variables for sensitive configuration
4. **Validation**: Always validate configuration values before use
5. **Access Control**: Limit access to configuration files and directories

## Contact

- **Security Email**: security@fork.vn
- **General Issues**: Use GitHub Issues for non-security bugs
- **Questions**: Use GitHub Discussions for questions

## Recognition

We appreciate security researchers and will acknowledge contributors in our security advisories (unless they prefer to remain anonymous).
