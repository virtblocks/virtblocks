FROM golang:1.12.5

ENV RUSTUP_ARCH=x86_64-unknown-linux-gnu \
    RUSTUP_VERSION=1.18.3 \
    RUSTUP_SHA256=a46fe67199b7bcbbde2dcbc23ae08db6f29883e260e23899a88b9073effc9076 \
    RUST_VERSION=1.37.0 \
    RUSTUP_HOME=/usr/local/rustup \
    CARGO_HOME=/usr/local/cargo \
    PATH=/usr/local/cargo/bin:$PATH

RUN apt-get update && \
    apt-get install -y libtool libtool-bin && \
    wget https://static.rust-lang.org/rustup/archive/$RUSTUP_VERSION/$RUSTUP_ARCH/rustup-init && \
    echo "$RUSTUP_SHA256 rustup-init" | sha256sum -c - && \
    chmod +x rustup-init && \
    ./rustup-init -y --no-modify-path --default-toolchain $RUST_VERSION && \
    rm rustup-init && \
    rustup component add rustfmt clippy
