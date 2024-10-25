import { Newsletter } from "../models";

interface NewsletterResponse {
  newsletters: Newsletter[];
}

export const getNewsletters = async (token: string): Promise<Newsletter[]> => {
  try {
    const response = await fetch(
      `${process.env.NEWSLETTER_DOCKER_SERVICE_API}/v1/newsletters`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      throw new Error("Error getting newsletters");
    }

    const data: NewsletterResponse = await response.json();
    return data.newsletters;
  } catch (e) {
    console.error("Error getting newsletters", e);
    return [];
  }
};
