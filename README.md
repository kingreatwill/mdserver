# README
README


## Start
`go env -w GOPROXY=https://goproxy.cn,direct`

`go mod tidy`

## demo

### katex
`$...$`
∪$\cup$`\cup`

`$$...$$`

$$
\cup
$$

$$
  \begin{align}
  D(x) &= \int_{x_0}^x P(x^{\prime})\,\mathrm{dx^{\prime}}  \\
  &= C\int_{x_0}^x x^{\prime n}\,\mathrm{dx^{\prime}} \\
  &= \frac{C}{n+1}(x^{n+1}-x_0^{n+1}) \\
  &\equiv  y
  \end{align}
$$

matrix

$$
\begin{matrix}
    1&x&x^2\\
    1&y&y^2\\
    1&z&z^2\\
\end{matrix}
$$



### mermaid

Sequence diagram

```mermaid
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    loop Healthcheck
        John->>John: Fight against hypochondria
    end
    Note right of John: Rational thoughts <br/>prevail!
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
```

Gantt diagram

```mermaid
gantt
dateFormat  YYYY-MM-DD
title Adding GANTT diagram to mermaid
excludes weekdays 2014-01-10

section A section
Completed task            :done,    des1, 2014-01-06,2014-01-08
Active task               :active,  des2, 2014-01-09, 3d
Future task               :         des3, after des2, 5d
Future task2               :         des4, after des3, 5d
```