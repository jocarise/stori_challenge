import { Recipient } from "../models";

export interface AddRecipientDto {
  email: string;
  newsletterId: string;
  token: string;
}

export const addRecipient = async ({
  email,
  newsletterId,
  token,
}: AddRecipientDto): Promise<Recipient | undefined> => {
  try {
    const response = await fetch(
      `${process.env.NEWSLETTER_SERVICE_API}/v1/newsletters/` + newsletterId,
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          email: email,
        }),
      }
    );

    if (!response.ok) {
      throw new Error("addRecipient failed");
    }

    const data: Recipient = await response.json();
    return data;
  } catch (e) {
    console.error("error addRecipient", e);
    return;
  }
};
