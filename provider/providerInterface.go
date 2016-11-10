package main

type ProviderInterface interface {
	GetInfo() (Information, error)
}
