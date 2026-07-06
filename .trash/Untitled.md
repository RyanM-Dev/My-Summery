
```mermaid
flowchart LR
    subgraph existing [Implemented]
        CP[CreatePayment]
        CT[CreateTransaction]
        UT[UpdateTransaction]
        GPU[GetPaymentByUUID]
        GPT[GetPaymentTransactionsByPaymentID]
        UP[UpdatePayment]
    end

    subgraph missing [Missing for planned flows]
        GTT[GetTransactionByGatewayToken]
        GPI[GetPaymentByID]
        GGC[GetGatewayByCode]
        GGI[GetGatewayByID]
        UPS[UpdatePaymentStatus guarded]
        TX[RunInTx]
        LCK[GetPaymentByIDForUpdate]
    end

    Initiate --> GGC --> CP --> CT --> UT
    Callback --> GTT --> LCK --> TX --> UT
    TX --> UPS
    VerifyRPC --> GPU
    VerifyRPC --> GPI
    ```