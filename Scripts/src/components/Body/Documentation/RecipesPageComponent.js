import React from "react"

export default function RecipesPageComponent() {
    return [
        {
            title: "Recipes",
            paragraph:
                <div>
                    <p>
                        <strong>Note:</strong> This is purely for developing FFXIV Profit. This endpoint is really just a proxy of accessing XIVAPI.
                        If you really wanted to get recipe information, go to <a href="https://xivapi.com/">XIVAPI</a> directly.
                    </p>
                    <p>
                        This is the recipe endpoint that just stores XIVAPI recipe data to our own database. This allows us to reduce load onto XIVAPI,
                        for repeated calls on the same recipe.
                    </p>
                </div>,
        },
        {
            title: "GET /recipe",
            paragraph:
                <div>
                    <p>
                        GET {window.location.protocol}//{window.location.hostname}/api/recipe/33180
                    </p>
                    <p>
                        Any recipe you need, you search for it using a specific recipe ID.
                        To obtain this, use a frontend framework to use XIVAPI's search, and filter it by recipe.
                    </p>
                    <p>
                        For example, {"https://xivapi.com/search?indexes=recipe&filters=&string=Aiming"},
                        would search for a recipe that has Aiming in it. Then it would return a payload back with the information about that recipe,
                        including the recipe ID.
                    </p>
                    <p>
                        Our frontend caches these search responses, and our backend stores the recipe info into our database.
                    </p>
                </div>,
        },
    ]
}