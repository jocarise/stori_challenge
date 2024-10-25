flowchart TD

    %% GetCategories
    subgraph GetCategories
        A[Start: Call GetCategories_FUNC] --> B[Retrieve categories from repository]
        B --> C[Return list of categories]
    end

    %% GetNewsletters
    subgraph GetNewsletters
        D[Start: Call GetNewsletters_FUNC] --> E[Retrieve newsletters from repository]
        E --> F[Return list of newsletters]
    end

    %% GetNewsletterById
    subgraph GetNewsletterById
        G[Start: Call GetNewsletterById_FUNC] --> H[Retrieve newsletter by ID from repository]
        H --> I{Is newsletter found?}
        I -- Yes --> J[Return found newsletter]
        I -- No --> K[Return nil and error]
    end

    %% UnsubscribeFromNewsletters
    subgraph UnsubscribeFromNewsletters
        L[Start: Call UnsubscribeFromNewsletters_FUNC] --> M[Remove recipient from all newsletters]
        M --> N[Return success or error]
    end

    %% UnsubscribeFromSpecificCategory
    subgraph UnsubscribeFromSpecificCategory
        O[Start: Call UnsubscribeFromSpecificCategory_FUNC] --> P[Remove recipient from specific category newsletters]
        P --> Q[Return success or error]
    end
