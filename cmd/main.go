package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/spf13/cobra"

	"github.com/8january/password-manager/internals/database"
)

var db = database.Init("local")

func main() {

	err := db.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	var rootCmd = &cobra.Command{
		Use:   "password-manager",
		Short: "Gerenciador de senhas simples em Go",
	}

	rootCmd.AddCommand(Save())
	rootCmd.AddCommand(Get())
	rootCmd.AddCommand(List())
	rootCmd.AddCommand(Delete())
	rootCmd.AddCommand(Update())
	rootCmd.AddCommand(Gen())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Erro ao executar a CLI: %v", err)
	}
}

func Save() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Salvar nova senha",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			service, _ := cmd.Flags().GetString("service")
			password, _ := cmd.Flags().GetString("password")
			passphrase, _ := cmd.Flags().GetString("passphrase")

			if name == "" || service == "" || password == "" || passphrase == "" {
				fmt.Println("Todos os campos são obrigatórios")
				return
			}

			db.Save(name, service, password, passphrase)
			fmt.Println("Senha salva com sucesso!")
		}}

	cmd.Flags().StringP("name", "n", "", "Nome da conta")
	cmd.Flags().StringP("service", "s", "", "Serviço da conta")
	cmd.Flags().StringP("password", "p", "", "Senha do serviço")
	cmd.Flags().StringP("passphrase", "k", "", "Chave mestra")

	return cmd

}

func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a password by ID",
		Run: func(cmd *cobra.Command, args []string) {

			id, _ := cmd.Flags().GetInt("id")
			passphrase, _ := cmd.Flags().GetString("passphrase")
			pwd := db.Get(id, passphrase)
			fmt.Println(pwd)

		},
	}

	cmd.Flags().String("id", "", "id da senha")
	cmd.Flags().StringP("passphrase", "k", "", "senha mestra")

	return cmd
}

func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all passwords",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listando todas as senhas")
			passwords := db.List()
			for _, p := range passwords {
				fmt.Printf(
					"%d | %s | %s | %s | %s\n",
					p.ID,
					p.Name,
					p.Service,
					p.Password,
					p.CreatedAt,
				)
			}
		},
	}
	return cmd
}

func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete password by id",
		Run: func(cmd *cobra.Command, args []string) {

			id, _ := cmd.Flags().GetInt("id")
			passphrase, _ := cmd.Flags().GetString("passphrase")
			db.Delete(id, passphrase)
		},
	}
	cmd.Flags().String("id", "", "id da senha")
	cmd.Flags().StringP("passphrase", "k", "", "senha mestra")
	return cmd
}

func Update() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update password by id",
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := cmd.Flags().GetInt("id")
			password, _ := cmd.Flags().GetString("password")
			oldPassphrase, _ := cmd.Flags().GetString("oldpassphrase")
			newPassphrase, _ := cmd.Flags().GetString("newpassphrase")

			db.Update(id, password, oldPassphrase, newPassphrase)
		},
	}
	cmd.Flags().String("id", "", "id da senha")
	cmd.Flags().StringP("password", "p", "", "nova senha")
	cmd.Flags().StringP("oldpassphrase", "o", "", "antiga senha mestra")
	cmd.Flags().StringP("newpassphrase", "n", "", "nova senha mestra")

	return cmd
}

func Gen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate a random password",
		Run: func(cmd *cobra.Command, args []string) {
			pLen, _ := cmd.Flags().GetInt("length")
			special, _ := cmd.Flags().GetBool("special")

			password := generatePassword(pLen, special)
			fmt.Println(password)
		},
	}
	cmd.Flags().IntP("length", "l", 16, "Length of the password")
	cmd.Flags().BoolP("special", "s", false, "Include special characters")
	return cmd
}

func generatePassword(length int, includeSpecial bool) string {
	// Character sets
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	// Create character pool based on parameters
	charPool := lowercase + uppercase + numbers
	if includeSpecial {
		charPool += special
	}

	// Generate random password
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charPool[rand.Intn(len(charPool))]
	}

	return string(password)
}
