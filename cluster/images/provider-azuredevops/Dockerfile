# Build stage for multi-architecture support
FROM alpine:3.21

# Install required packages
RUN apk --no-cache add ca-certificates bash

# Build arguments for multi-arch support
ARG TARGETOS
ARG TARGETARCH

# Add the provider binary
ADD bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/provider

# User configuration (non-root)
ENV USER_ID=65532

# Terraform configuration arguments
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX

# Terraform environment setup
ENV PLUGIN_DIR=/terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0

RUN mkdir -p ${PLUGIN_DIR}

# Download and install Terraform
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/terraform.zip

# Download Terraform provider
ADD ${TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp/provider.zip

# Copy Terraform configuration
COPY terraformrc.hcl ${TF_CLI_CONFIG_FILE}

# Install Terraform and provider
RUN unzip /tmp/terraform.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform.zip \
  && unzip /tmp/provider.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/provider.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform

# Set runtime environment variables for provider controller
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

# Run as non-root user
USER ${USER_ID}

# Expose metrics port
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["provider"]
