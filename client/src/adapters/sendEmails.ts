export const sendEmails = async (
  token: string,
  newsletterId: string
): Promise<number> => {
  try {
    const response = await fetch(
      `${process.env.NEWSLETTER_SERVICE_API}/v1/newsletters/${newsletterId}/send`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      throw new Error("sendEmails failed");
    }

    return response.status;
  } catch (e) {
    console.error("error sendEmails", e);
    return 500;
  }
};
