package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SolAddress struct {
	Address    string  `json:"address"`
	Nickname   *string `json:"nickname"`
	SOLBalance string  `json:"sol_balance"`
	Updated    int64   `json:"updated"`
}

type Handler struct {
	Addresses map[string]*SolAddress
}

func New() *Handler {
	return &Handler{
		Addresses: make(map[string]*SolAddress),
	}
}

func (h *Handler) SubmitAddress(c *fiber.Ctx) error {
	req := new(SubmitAddressReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if req.SolAddress == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid address",
		})
	}

	solBalance, err := GetSolBalance(req.SolAddress)
	fmt.Printf("solBalance: %s\n", solBalance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	now := time.Now().Unix()
	newAddress := false
	if _, ok := h.Addresses[req.SolAddress]; ok {
		e := h.Addresses[req.SolAddress]
		e.SOLBalance = solBalance
		e.Updated = now
	} else {
		newAddress = true
		h.Addresses[req.SolAddress] = &SolAddress{
			Address:    req.SolAddress,
			SOLBalance: solBalance,
			Updated:    now,
		}
	}

	res := SubmitAddressRes{
		NewAddress: newAddress,
	}

	return c.JSON(res)
}

func (h *Handler) GetAddresses(c *fiber.Ctx) error {
	rateData, err := GetExchangeRateInfo()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	solRate, err := strconv.ParseFloat(rateData.Data.Rates["SOL"], 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	addresses := []*SolAddress{}
	for _, a := range h.Addresses {
		addresses = append(addresses, a)
	}

	res := GetAddressesRes{
		Addresses:    addresses,
		ExchangeRate: 1 / solRate,
	}

	return c.JSON(res)
}

func (h *Handler) NameAddress(c *fiber.Ctx) error {
	req := new(NameAddressReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if _, ok := h.Addresses[req.SolAddress]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Address not found",
		})
	}

	e := h.Addresses[req.SolAddress]
	e.Nickname = &req.Nickname

	return c.Status(200).SendString("Done")
}

func (h *Handler) ClearAddresses(c *fiber.Ctx) error {
	h.Addresses = make(map[string]*SolAddress)
	return c.Status(200).SendString("Done")
}
