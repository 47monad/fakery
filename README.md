# Fakery

[![Go Reference](https://pkg.go.dev/badge/github.com/47monad/fakery.svg)](https://pkg.go.dev/github.com/47monad/fakery)
[![Go Report Card](https://goreportcard.com/badge/github.com/47monad/fakery)](https://goreportcard.com/report/github.com/47monad/fakery)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/47monad/fakery/workflows/CI/badge.svg)](https://github.com/47monad/fakery/actions)
[![Coverage Status](https://coveralls.io/repos/github/47monad/fakery/badge.svg?branch=main)](https://coveralls.io/github/47monad/fakery?branch=main)

> ⚠️ **Under Active Development** - This library is currently under active development. APIs may change before the first stable release.

A modern, type-safe Go library for generating realistic fake data with comprehensive internationalization support. Fakery provides a clean, functional API with flexible builder patterns for generating test data, seeding databases, and creating realistic datasets.

## Features

- 🌍 **Internationalization Support** - Generate locale-specific data with automatic fallbacks
- 🔧 **Type-Safe Generics** - Leverage Go's type system for better developer experience
- 🏗️ **Flexible Builder Pattern** - Configure data generation with intuitive options
- 🌐 **Internet Data Generation** - Domains, emails, IP addresses, browser info, and more
- 👥 **Person Data** - Names with gender and language support
- 🏢 **Company Data** - Business names and buzzwords
- 📱 **Communication Data** - Phone numbers and contact information
- 📝 **Lorem Ipsum** - Text generation with word count controls
- 🎯 **Dual API** - Use standalone functions or instance-based generators

## Installation

```bash
go get github.com/47monad/fakery
```

## Quick Start

### Standalone Functions

```go
package main

import (
    "fmt"
    "github.com/47monad/fakery/fk/fkinternet"
    "github.com/47monad/fakery/fk/fkperson"
)

func main() {
    // Generate internet data
    domain := fkinternet.Domain()
    email := fkinternet.Email()
    ipv4 := fkinternet.IPv4()
    ipv6 := fkinternet.IPv6()
    
    fmt.Printf("Domain: %s\n", domain)      // example: buzztech.com
    fmt.Printf("Email: %s\n", email)        // example: john.doe@innovate.org
    fmt.Printf("IPv4: %s\n", ipv4)          // example: 192.168.1.100
    fmt.Printf("IPv6: %s\n", ipv6)          // example: 2001:db8::1
    
    // Generate person data
    firstName := fkperson.FirstName()
    lastName := fkperson.LastName()
    
    fmt.Printf("Name: %s %s\n", firstName, lastName)
}
```

### Instance-Based Generation

```go
package main

import (
    "fmt"
    "github.com/47monad/fakery/fk/fkinternet"
    "github.com/47monad/fakery/fk/fkopts"
)

func main() {
    // Create a faker instance with custom options
    faker := fkinternet.New(
        fkopts.Faker[fkdata.Internet]().SetLang("en"),
    )
    
    // Use the instance for consistent generation
    domain := faker.Domain()
    email := faker.Email()
    
    fmt.Printf("Domain: %s\n", domain)
    fmt.Printf("Email: %s\n", email)
}
```

## Internet Data Generation

The `fkinternet` package provides comprehensive internet-related fake data:

```go
import "github.com/47monad/fakery/fk/fkinternet"

// Network addresses
ipv4 := fkinternet.IPv4()          // 192.168.1.42
ipv6 := fkinternet.IPv6()          // 2001:db8:85a3::8a2e:370:7334
mac := fkinternet.MacAddr()        // aa:bb:cc:dd:ee:ff

// Web-related data
domain := fkinternet.Domain()      // techbuzz.com
email := fkinternet.Email()        // jane.smith@innovate.org
username := fkinternet.Username()  // wilson123

// Browser information
browser := fkinternet.Browser()         // Chrome, Firefox, Safari
engine := fkinternet.BrowserEngine()    // WebKit, Gecko, Blink
```

## Options and Configuration

Fakery uses a flexible builder pattern for configuration:

```go
import (
    "github.com/47monad/fakery/fk/fkperson"
    "github.com/47monad/fakery/fk/fkopts"
)

// Configure name generation
firstName := fkperson.FirstName(
    fkopts.Name().
        SetLang("en").
        SetGender("female"),
)

lastName := fkperson.LastName(
    fkopts.LastName().SetLang("en"),
)
```

## Internationalization

Fakery supports multiple languages with automatic fallbacks:

```go
// Generate German names
germanFirst := fkperson.FirstName(
    fkopts.Name().SetLang("de"),
)

// Generate Spanish names  
spanishLast := fkperson.LastName(
    fkopts.LastName().SetLang("es"),
)

// If locale data isn't available, automatically falls back to English
```

## Architecture

Fakery is built with several key design principles:

- **Type Safety**: Heavy use of Go generics for compile-time safety
- **Modularity**: Separate packages for different data domains
- **Extensibility**: Easy to add new data types and locales
- **Performance**: Efficient data binding and sampling mechanisms
- **Testability**: Clean interfaces and dependency injection

### Core Components

- **`fk/`** - Core context and configuration management
- **`fkinternet/`** - Internet-related data generation
- **`fkperson/`** - Person names and demographic data
- **`fkcompany/`** - Business and company data
- **`fkopts/`** - Builder pattern implementations for all options
- **`internal/`** - Internal utilities for randomization, sampling, and templating

## Data Sources

Fakery uses JSON-based locale files for realistic, culturally-appropriate data generation. The data binding system automatically handles:

- Language-specific data loading
- Fallback to English when locale data is unavailable  
- Type-safe data structures with generics
- Efficient sampling from data pools

## Contributing

This project is under active development. Contributions are welcome! Please feel free to:

- Report bugs and issues
- Suggest new features or data types
- Submit pull requests
- Add new locale data

## Roadmap

- [ ] Additional data generators (addresses, credit cards, dates)
- [ ] Enhanced browser and user agent generation
- [ ] Custom data providers
- [ ] Performance optimizations
- [ ] Comprehensive test coverage
- [ ] Documentation improvements

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Related Projects

If you're looking for fake data generation in other languages:
- **JavaScript**: [Faker.js](https://github.com/faker-js/faker)
- **Python**: [Faker](https://github.com/joke2k/faker)
- **Ruby**: [FFaker](https://github.com/ffaker/ffaker)
- **Java**: [JavaFaker](https://github.com/DiUS/java-faker)

---

*Built with ❤️ using Go generics and modern design patterns in 47monad*
