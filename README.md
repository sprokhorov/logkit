# logkit
LogKit provides a lightweight, standardized logger interface, designed to streamline the process of switching or integrating different logging implementations with minimal friction.

## Overview
Logging is a crucial aspect of any application, providing visibility into the application's behavior and helping with debugging and monitoring. However, logging implementations can vary significantly between different libraries and frameworks. LogKit addresses this challenge by offering a simple and consistent logger interface that can be implemented using various logging libraries or custom loggers.

With LogKit, you can easily swap out the underlying logging mechanism without altering the core application code. This flexibility ensures that you can adopt the best logging tool for your needs while maintaining a consistent interface across your project.

## Features
* **Standardized Interface:** Provides a unified logging interface that can be implemented by any logging library or custom solution.
* **Easy Integration:** Integrate LogKit into existing projects with minimal code changes.
* **Extensible:** Allows developers to create custom loggers that adhere to the LogKit interface.
* **Lightweight:** Minimalistic design with no external dependencies.


## Interface Definition
The Logger interface defines the following methods for logging at various levels of severity:
```
type Logger interface {
    Debug(args ...interface{})
    Debugf(format string, args ...interface{})
    Info(args ...interface{})
    Infof(format string, args ...interface{})
    Warn(args ...interface{})
    Warnf(format string, args ...interface{})
    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    Fatal(args ...interface{})
    Fatalf(format string, args ...interface{})
}
```

## Default Implementation
LogKit comes with a default implementation called DefaultLogger, which uses Go's standard log package. This implementation can be used out-of-the-box, but it can also serve as a reference for creating custom loggers.
```
type DefaultLogger struct{}

func (dl *DefaultLogger) Debug(args ...interface{}) {
    log.Println(args...)
}

func (dl *DefaultLogger) Debugf(format string, args ...interface{}) {
    log.Printf(format, args...)
}

func (dl *DefaultLogger) Info(args ...interface{}) {
    log.Println(args...)
}

func (dl *DefaultLogger) Infof(format string, args ...interface{}) {
    log.Printf(format, args...)
}

func (dl *DefaultLogger) Warn(args ...interface{}) {
    log.Println(args...)
}

func (dl *DefaultLogger) Warnf(format string, args ...interface{}) {
    log.Printf(format, args...)
}

func (dl *DefaultLogger) Error(args ...interface{}) {
    log.Println(args...)
}

func (dl *DefaultLogger) Errorf(format string, args ...interface{}) {
    log.Printf(format, args...)
}

func (dl *DefaultLogger) Fatal(args ...interface{}) {
    log.Fatal(args...)
}

func (dl *DefaultLogger) Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}
```

## Example Usage
Below is a simple example of how to use LogKit in your Go application:
```
package main

import (
    "github.com/yourusername/logkit"
)

func main() {
    var logger logkit.Logger = &logkit.DefaultLogger{}

    logger.Debug("This is a debug message")
    logger.Infof("This is an %s message with formatting", "info")
    logger.Warn("This is a warning")
    logger.Error("This is an error")
    // logger.Fatal("This is a fatal error") // Uncommenting this will terminate the program

    // Switching to a custom logger
    logger = &MyCustomLogger{}
    logger.Info("Now using a custom logger!")
}

// MyCustomLogger is an example of a custom logger implementing the LogKit interface
type MyCustomLogger struct{}

func (c *MyCustomLogger) Debug(args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Debugf(format string, args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Info(args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Infof(format string, args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Warn(args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Warnf(format string, args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Error(args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Errorf(format string, args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Fatal(args ...interface{}) {
    // Custom implementation
}

func (c *MyCustomLogger) Fatalf(format string, args ...interface{}) {
    // Custom implementation
}
```

### Explanation
* **DefaultLogger:** By default, you can use DefaultLogger, which utilizes the Go standard log package.
* **CustomLogger:** You can also implement your own custom logger by creating a type that satisfies the Logger interface.

## Getting Started
To install the LogKit package, use:
```
go get github.com/yourusername/logkit
```

Then import it into your project:
```
import "github.com/yourusername/logkit"
```

You can start using the Logger interface right away or implement your own custom logger that conforms to the interface.
