package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_user/pkg/clients/rabbitmq"
	inter "github.com/shivaraj-shanthaiah/code_orbit_user/pkg/repo/interfaces"
	"gopkg.in/robfig/cron.v2"
)

type CronInstance struct {
	cron *cron.Cron
	Repo inter.UserRepoInter
}

func NewCron(repo inter.UserRepoInter) *CronInstance {
	return &CronInstance{
		cron: cron.New(),
		Repo: repo,
	}
}

func (c *CronInstance) InitCron() {
    log.Println("Starting cron scheduler...")
    c.cron.Start()
    log.Println("Cron scheduler started.")
}

func (c *CronInstance) SheduleJob() error {
	_, err := c.cron.AddFunc("10 14 * * *", func() {
		c.CheckMembership()
	})
	return err
}

func (c *CronInstance) CheckMembership() {
	log.Println("Checking membership expiry for users...")
	users, err := c.Repo.FindAllPrimeUsers()
	if err != nil {
		log.Printf("Failed to retrieve prime users: %v", err)
		return
	}

	for _, user := range users {
		timeRemaining := time.Until(user.MembershipExpiryDate).Hours()

		// If the membership expires in less than 48 hours, send a warning email
		if timeRemaining > 0 && timeRemaining <= 800 {
			subject := "Membership Expiry Warning"
			message := fmt.Sprintf("Dear %s, your membership will expire in 2 days. Please renew it to continue enjoying prime benefits.", user.UserName)

			if err := rabbitmq.SendMail(user.Email, message, subject); err != nil {
				log.Printf("Failed to send warning email to user %s: %v", user.Email, err)
			} else {
				log.Printf("Sent warning email to user %s", user.Email)
			}
		}

		// If the membership has already expired, send a notification and update the status
		if timeRemaining <= 0 {
			user.IsPrimeMember = false
			if err := c.Repo.UpdateUserMembership(user); err != nil {
				log.Printf("Failed to update user membership status for %s: %v", user.Email, err)
				continue
			}

			subject := "Membership Expired"
			message := fmt.Sprintf("Dear %s, your membership has expired. Please renew it to regain access to prime features.", user.UserName)

			if err := rabbitmq.SendMail(user.Email, message, subject); err != nil {
				log.Printf("Failed to send expiration email to user %s: %v", user.Email, err)
			} else {
				log.Printf("Sent expiration email to user %s", user.Email)
			}
		}
	}
}
