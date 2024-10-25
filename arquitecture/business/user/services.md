flowchart TD
    subgraph AuthUserFlow
        A[Start: Call AuthUser_FUNC] --> B[Validate email format]
        B -->|Invalid| C[Return '', error 'invalid email or password']
        B -->|Valid| D[Validate password format]
        
        D -->|Invalid| C
        D -->|Valid| E[Fetch user by email from repo]
        
        E --> F{User found?}
        F -- No --> G[Return '', error 'invalid email or password']
        
        F -- Yes --> H[Compare provided password with stored password]
        
        H -->|Mismatch| I[Return '', error 'invalid email or password']
        H -->|Match| J[Generate JWT token]
        
        J --> K{Token generation successful?}
        K -- No --> L[Return '', error 'invalid email or password']
        
        K -- Yes --> M[Return token]
    end

    subgraph CreateUserFlow
        N[Start: Call CreateUser_FUNC] --> O[Validate UUID]
        O -->|Invalid| P[Return error 'invalid UUID']
        
        O -->|Valid| Q[Validate email format]
        Q -->|Invalid| R[Return error 'invalid email']
        
        Q -->|Valid| S[Validate password format]
        S -->|Invalid| T[Return error 'invalid password']
        
        S -->|Valid| U[Hash password]
        
        U --> V{Hashing successful?}
        V -- No --> W[Return error 'invalid password']
        
        V -- Yes --> X[Set hashed password on user]
        X --> Y[Create user in repo]
        
        Y --> Z{Creation successful?}
        Z -- No --> AA[Return error]
        
        Z -- Yes --> AB[Return nil]
    end