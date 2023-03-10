terraform {
  backend "azurerm" {
    resource_group_name  = "terraform-state"
    storage_account_name = "terraformstate0419"
    container_name       = "shared-tfstate"
    key                  = "prod.terraform.tfstate"
  }

  required_version = "1.2.8"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.20.0"
    }
    azuread = {
      source = "hashicorp/azuread"
      version = "2.28.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "random_pet" "suffix" {
  length    = 1
  separator = ""
}

data "azurerm_client_config" "current" {}

data "azurerm_subscription" "current" {}

resource "azurerm_resource_group" "this" {
  name     = "e2e-shared"
  location = "West Europe"
}
