package repository

type App interface {
	Run(args []string) error
}
