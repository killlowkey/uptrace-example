## 日志记录

Q1: 在 golang 中分为 controller、service、store 三层，如果 store 发生错误，应该在那一层记录日志

A: 在 Go 语言的典型三层架构中——controller、service、store——处理错误和记录日志的最佳实践通常遵循以下原则：

1. Controller 层：处理来自外部的请求和响应。在 controller
   层记录日志通常是首选做法，因为在这里可以捕获和记录关于错误的完整上下文信息，包括来自请求的数据、错误的性质以及可能影响到的用户或系统操作。将日志记录集中在
   controller 层也有助于保持代码的整洁和一致性。
   在 Go 语言的三层架构中，通常建议在 controller 层进行错误处理和日志记录，因为这样可以提供最全面的上下文信息。然而，根据具体的应用需求和架构设计，service
   层在某些情况下也可能参与日志记录。
2. Service 层：通常处理业务逻辑。当 service 层从 store 层接收到错误时，它可以决定如何处理这些错误。在某些情况下，service
   层可能会记录日志，尤其是当错误涉及到业务逻辑决策或需要进一步的上下文信息时。然而，通常建议将错误传递给上层（controller
   层），在那里进行更统一的错误处理和日志记录。
3. Store 层：负责数据存储和检索。当错误发生时，通常应该将错误返回给调用它的 service 层，而不是在这里记录日志。Store
   层应保持其专注于数据存取的职责，避免承担日志记录等额外职责。

在 Go 语言的三层架构中，通常建议在 controller 层进行错误处理和日志记录，因为这样可以提供最全面的上下文信息。然而，根据具体的应用需求和架构设计，service 层在某些情况下也可能参与日志记录。


Q2: 应该记录什么类型日志

A:
1. 错误日志（Error Logs）：当发生错误时，无论是由于外部因素还是内部问题，都应记录详细的错误信息。这包括错误消息、堆栈跟踪（对于严重错误）、以及导致错误的任何相关数据（但注意避免记录敏感信息）。
2. 警告日志（Warning Logs）：用于记录可能指示潜在问题的非严重事件。例如，非预期的输入值、性能问题、或者预期发生的错误（例如，用户输入了无效的数据）。
3. 信息日志（Info Logs）：记录应用程序的一般运行信息，例如启动或关闭、配置更改、连接到数据库等。
4. 调试日志（Debug Logs）：这些日志提供更详细的信息，有助于开发人员调试代码。通常包括变量值、状态信息、函数调用顺序等。这类日志在生产环境中通常是关闭的，因为可能会生成大量数据。
5. 审计日志（Audit Logs）：用于记录关键业务流程和决策。例如，用户的认证和授权活动、关键数据的访问和修改记录等。这类日志对于满足合规性要求和进行安全分析非常重要。

当记录日志时，应该考虑以下几点：
1. 上下文信息：记录足够的信息以便理解日志条目的背景。例如，在哪个模块发生的、影响的用户或事务等。
2. 格式一致性：保持日志的格式一致，以便于日志解析和监控工具的使用。
3. 敏感信息处理：避免记录敏感信息，如密码、个人身份信息等。
4. 性能影响：考虑日志记录对应用性能的影响。大量的日志记录或非常频繁的日志记录可能对性能有负面影响。
5. 日志级别管理：合理使用日志级别（如 DEBUG、INFO、WARN、ERROR），并在不同的环境（如开发、测试、生产）中设置适当的日志级别。