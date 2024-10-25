import { Newsletter } from "../models";

interface CreateNewsletterDto {
  formData: FormData;
  token: string;
}

export const createNewsletter = async ({
  formData,
  token,
}: CreateNewsletterDto): Promise<Newsletter | undefined> => {
  try {
    const response = await fetch(
      `${process.env.NEWSLETTER_SERVICE_API}/v1/newsletters`,
      {
        method: "POST",
        body: formData,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      throw new Error("Error creating newsletter");
    }

    const result = await response.json();

    return result;
  } catch (e) {
    console.error("Error sending newsletter:", e);
    return;
  }
};
