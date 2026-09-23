# Robot Market - User Flow

Source: `Robot Market — Product Requirements Document (1).md`

This document describes the main user flows for Robot Market using Mermaid diagrams. The flows focus on the MVP/current scope and mark payment and wallet behavior as later-phase items.

---

## 1. Role Access Overview

```mermaid
flowchart TD
    Visitor[Visitor]
    Customer[Customer]
    Admin[Admin]
    Owner[Owner]

    Visitor --> PublicCatalog[Browse product listing]
    Visitor --> ProductDetails[View product details, images, specs, prices]
    Visitor --> PublicPages[View public informational pages]
    Visitor --> PurchaseIntent{Wants to add to cart or order?}
    PurchaseIntent --> AuthRequired[Register / login with mobile OTP]
    AuthRequired --> Customer

    Customer --> CustomerCart[Maintain persistent cart]
    Customer --> ConfigureProduct[Configure product modules]
    Customer --> SubmitOrder[Submit order]
    Customer --> CustomerDashboard[View dashboard, active orders, history]
    Customer --> CancelBeforeApproval[Cancel eligible orders before admin approval]
    Customer --> ReviewInvoice[Review final invoice / order]

    Admin --> CatalogManagement[Manage products, specs, images, modules, prices]
    Admin --> CustomerManagement[Find or register customers]
    Admin --> AssistedOrdering[Create orders for customers]
    Admin --> ReviewOrders[Review, modify, approve, decline orders]
    Admin --> InternalNotes[Add internal notes]

    Owner --> AdminManagement[Add or remove admins]
    Owner --> Admin
```

---

## 2. Customer Self-Service Ordering Flow

```mermaid
flowchart TD
    Start([Visitor opens Robot Market])
    Browse[Browse products]
    Details[View vending machine details]
    Select[Select vending machine]
    WantsCart{Add to cart / start order?}
    Auth{Authenticated?}
    OTP[Register or login using mobile OTP]
    AddToCart[Add product to persistent cart]
    Configure[Select compatible available modules]
    Price[Show calculated product + module price]
    Info{Required order info complete?}
    CompleteInfo[Enter or update customer and delivery info]
    Submit[Submit order]
    Submitted[Order status: Submitted]
    AdminReview[Admin review begins]
    UnderReview[Order status: Under Review]
    AdminDecision{Admin decision}
    RequestInfo[Request additional customer information]
    Modify[Admin modifies modules, configuration, price, or details]
    Decline[Order status: Declined]
    Approve[Order status: Approved]
    AwaitingPayment[Order status: Awaiting Payment]
    FinalReview[Customer reviews final invoice / order]
    Accept{Customer accepts final order?}
    PaymentLater[Payment gateway / wallet later phase]
    Cancel[Customer cancels before approval]
    Cancelled[Order status: Cancelled]
    End([Flow ends])

    Start --> Browse --> Details --> Select --> WantsCart
    WantsCart -- No --> Browse
    WantsCart -- Yes --> Auth
    Auth -- No --> OTP --> AddToCart
    Auth -- Yes --> AddToCart
    AddToCart --> Configure --> Price --> Info
    Info -- No --> CompleteInfo --> Info
    Info -- Yes --> Submit --> Submitted --> AdminReview --> UnderReview --> AdminDecision

    Submitted -. Customer may cancel before approval .-> Cancel --> Cancelled --> End
    UnderReview -. Customer self-editing becomes admin-controlled .-> AdminDecision

    AdminDecision -- Request info --> RequestInfo --> CompleteInfo
    AdminDecision -- Modify --> Modify --> AdminDecision
    AdminDecision -- Decline --> Decline --> End
    AdminDecision -- Approve --> Approve --> AwaitingPayment --> FinalReview --> Accept
    Accept -- Yes --> PaymentLater --> End
    Accept -- No --> End
```

---

## 3. Admin-Assisted Customer Registration and Ordering Flow

```mermaid
flowchart TD
    Start([Admin opens admin panel])
    Auth[Admin authenticates]
    FindCustomer[Search for existing customer]
    CustomerFound{Customer exists?}
    RegisterCustomer[Register customer with first name, last name, mobile number]
    AuditAccount[Record admin-created account audit event]
    SelectCustomer[Select customer account]
    CreateOrder[Create order on behalf of customer]
    SelectProduct[Select vending machine]
    ConfigureModules[Configure modules]
    EnterInfo[Enter required customer and delivery information]
    ReviewDraft[Review draft order]
    NeedChanges{Changes needed?}
    ModifyOrder[Modify configuration, modules, price, or details]
    Approve[Approve order]
    AuditOrder[Record admin-created / admin-approved order audit events]
    CustomerReview[Customer reviews final invoice / order]
    CustomerAccept{Customer accepts final order?}
    PaymentLater[Payment gateway / wallet later phase]
    End([Flow ends])

    Start --> Auth --> FindCustomer --> CustomerFound
    CustomerFound -- No --> RegisterCustomer --> AuditAccount --> SelectCustomer
    CustomerFound -- Yes --> SelectCustomer
    SelectCustomer --> CreateOrder --> SelectProduct --> ConfigureModules --> EnterInfo --> ReviewDraft --> NeedChanges
    NeedChanges -- Yes --> ModifyOrder --> ReviewDraft
    NeedChanges -- No --> Approve --> AuditOrder --> CustomerReview --> CustomerAccept
    CustomerAccept -- Yes --> PaymentLater --> End
    CustomerAccept -- No --> End
```

---

## 4. Order Lifecycle

```mermaid
stateDiagram-v2
    [*] --> Draft: Cart / order being prepared
    Draft --> Submitted: Customer submits order
    Draft --> Submitted: Admin creates order for customer
    Draft --> Cancelled: Customer cancels before submission

    Submitted --> Cancelled: Customer cancels before admin approval
    Submitted --> UnderReview: Admin starts review

    UnderReview --> UnderReview: Admin modifies details, modules, or price
    UnderReview --> Submitted: Admin requests more information
    UnderReview --> Declined: Admin declines order
    UnderReview --> Approved: Admin approves order

    Approved --> AwaitingPayment: Final invoice is available
    AwaitingPayment --> Paid: Payment succeeds in later phase
    AwaitingPayment --> Cancelled: Administrative cancellation only

    Declined --> [*]
    Cancelled --> [*]
    Paid --> [*]
```

---

## 5. Admin Catalog and Pricing Management Flow

```mermaid
flowchart TD
    Start([Admin opens admin panel])
    Auth[Admin authenticates]
    ChooseArea{Management area}

    ProductCRUD[Create / edit / deactivate products]
    Specs[Manage product specifications]
    Images[Manage product images]
    BasePrice[Manage base price]
    Visibility[Set product active / inactive visibility]

    ModuleCRUD[Create / edit / deactivate modules]
    ModulePrice[Manage module prices]
    ModuleAvailability[Manage module availability]
    Compatibility[Maintain future-ready compatibility data]

    Relationship[Associate modules with products]
    PublicCatalog[Public catalog reflects active products and available modules]
    Audit[Record important admin changes for auditability]
    End([Flow ends])

    Start --> Auth --> ChooseArea

    ChooseArea -- Products --> ProductCRUD --> Specs --> Images --> BasePrice --> Visibility --> Relationship
    ChooseArea -- Modules --> ModuleCRUD --> ModulePrice --> ModuleAvailability --> Compatibility --> Relationship

    Relationship --> Audit --> PublicCatalog --> End
```

---

## 6. Final Invoice Acceptance Rule

```mermaid
flowchart TD
    AdminReview[Admin reviews submitted or assisted order]
    Changes{Admin changes order?}
    NoChange[Keep customer-selected configuration and calculated price]
    Changed[Modify configuration, modules, price, or details]
    Approved[Approve final order]
    Invoice[Customer sees final invoice / order]
    PaymentIntent{Customer proceeds to payment?}
    Accepted[Final configuration and price accepted]
    NotAccepted[No payment acceptance recorded]

    AdminReview --> Changes
    Changes -- No --> NoChange --> Approved
    Changes -- Yes --> Changed --> Approved
    Approved --> Invoice --> PaymentIntent
    PaymentIntent -- Yes --> Accepted
    PaymentIntent -- No --> NotAccepted
```

