package cli

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/StackGuardian/sg-sdk-go/billing"
	"github.com/StackGuardian/sg-sdk-go/client"
)

func (a *app) billingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billing",
		Short: "Inspect and manage the organization's billing",
	}

	balance := a.command("balance", "Show the credit balances", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Billing.ListBalance(ctx, a.org)
			return err
		})

	dashboard := a.command("dashboard-url", "Get the URL of the embeddable billing dashboard", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := billing.DashboardRequest{Type: flagString(cmd, "type")}
			if req.Type != nil {
				if err := checkEnum("type", *req.Type, "invoices", "usage"); err != nil {
					return err
				}
			}
			_, err := c.Billing.GetDashboardUrl(ctx, a.org, &req)
			return err
		})
	dashboard.Flags().String("type", "", "dashboard type (invoices, usage)")

	invoices := a.command("invoices", "List invoices", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := billing.ListInvoicesRequest{Limit: flagInt(cmd, "limit"), NextPage: flagString(cmd, "next-page")}
			_, err := c.Billing.ListInvoices(ctx, a.org, &req)
			return err
		})
	invoices.Flags().Int("limit", 0, "page size (1-100)")
	invoices.Flags().String("next-page", "", "cursor from the previous response")

	details := a.command("details", "Show the current package and billing information", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Billing.GetBillingDetails(ctx, a.org)
			return err
		})

	changePlan := a.command("change-plan", "Move the organization to another billing package", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			req := billing.ChangePlanRequest{
				TargetPackageId: flagString(cmd, "target-package-id"),
				TargetPackage:   flagString(cmd, "target-package"),
			}
			_, err := c.Billing.ChangePlan(ctx, a.org, &req)
			return err
		})
	changePlan.Flags().String("target-package-id", "", "id of the target package")
	changePlan.Flags().String("target-package", "", "name of the target package, e.g. \"SG | Standard\"")
	changePlan.MarkFlagsOneRequired("target-package-id", "target-package")
	changePlan.MarkFlagsMutuallyExclusive("target-package-id", "target-package")

	profile := &cobra.Command{Use: "profile", Short: "Manage the Stripe billing profile"}
	profileSetup := a.command("setup", "Create the billing profile (once per organization)", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req billing.BillingProfileRequest
			if err := readOptionalBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Billing.SetupBillingProfile(ctx, a.org, &req)
			return err
		})
	addBodyFlags(profileSetup)
	profileUpdate := a.command("update", "Update the billing profile's email, name or address", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, cmd *cobra.Command, _ []string) error {
			var req billing.BillingProfileRequest
			if err := readBody(cmd, &req); err != nil {
				return err
			}
			_, err := c.Billing.UpdateBillingProfile(ctx, a.org, &req)
			return err
		})
	addBodyFlags(profileUpdate)
	profile.AddCommand(profileSetup, profileUpdate)

	payment := &cobra.Command{Use: "payment-methods", Short: "Manage the cards on the billing profile"}
	paymentList := a.command("list", "List payment methods", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Billing.ListPaymentMethods(ctx, a.org)
			return err
		})
	paymentIntent := a.command("setup-intent", "Create a Stripe SetupIntent for adding a card", cobra.NoArgs,
		func(ctx context.Context, c *client.Client, _ *cobra.Command, _ []string) error {
			_, err := c.Billing.CreatePaymentSetupIntent(ctx, a.org)
			return err
		})
	paymentDefault := a.command("set-default PAYMENT_METHOD_ID", "Make a card the default payment method", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Billing.SetDefaultPaymentMethod(ctx, a.org, &billing.PaymentMethodRequest{PaymentMethodId: args[0]})
			return err
		})
	paymentDetach := a.command("detach PAYMENT_METHOD_ID", "Remove a card", cobra.ExactArgs(1),
		func(ctx context.Context, c *client.Client, _ *cobra.Command, args []string) error {
			_, err := c.Billing.DetachCard(ctx, a.org, &billing.PaymentMethodRequest{PaymentMethodId: args[0]})
			return err
		})
	payment.AddCommand(paymentList, paymentIntent, paymentDefault, paymentDetach)

	cmd.AddCommand(balance, dashboard, invoices, details, changePlan, profile, payment)
	return cmd
}
