import { Category } from "../models";

interface CategoriesResponse {
  categories: Category[];
}

export const getCategories = async (token: string): Promise<Category[]> => {
  try {
    const response = await fetch(
      `${process.env.NEWSLETTER_DOCKER_SERVICE_API}/v1/newsletters/categories`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      throw new Error("Error getting categories");
    }

    const data: CategoriesResponse = await response.json();
    return data.categories;
  } catch (e) {
    console.error("Error getting categories", e);
    return [];
  }
};
