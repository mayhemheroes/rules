FROM golang:1.18 as builder

COPY . /rules
WORKDIR /rules/mayhem

RUN go mod tidy

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzRules.a && \
    clang -fsanitize=fuzzer fuzzRules.a -o fuzzRules.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /rules/mayhem/fuzzRules.libfuzzer /