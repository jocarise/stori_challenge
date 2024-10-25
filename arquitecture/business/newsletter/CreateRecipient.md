flowchart TD
    A[Start: Call CreateRecipient_FUNC] --> B[Retrieve baseURL]
    B --> C{Validate email format}
    C -- Invalid --> D[Return nil, error 'invalid email format']
    C -- Valid --> E[Find recipient by email]
    
    E --> F{Does recipient exist?}
    F -- Yes --> G[Associate existing recipient with newsletter]
    G --> H{Was association successful?}
    H -- Yes --> I[Return existing recipient]
    H -- No --> J[Return nil, error 'internal server error']

    F -- No --> K[Retrieve newsletter by ID]
    K --> L{Was newsletter found?}
    L -- No --> M[Return nil, error 'internal server error']
    
    L -- Yes --> N[Create new recipient]
    N --> O[Save new recipient to repository]
    O --> P{Was saving successful?}
    P -- No --> Q[Return nil, error 'internal server error']
    
    P -- Yes --> R[Associate new recipient with newsletter]
    R --> S{Was association successful?}
    S -- No --> T[Return nil, error 'internal server error']
    S -- Yes --> U[Return new recipient]