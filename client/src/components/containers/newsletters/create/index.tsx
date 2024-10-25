import { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import { getCookie } from "cookies-next";
import { useRouter } from "next/router";
import { Category } from "../../../../models";
import { FormValues } from "../../../../schemas/newsletter";
import { createNewsletter } from "../../../../adapters/createNewsletter";
import { NewsletterForm } from "./form";

interface CreateNewsletterContainerProps {
  categories: Category[];
}

export const CreateNewsletterContainer = ({
  categories,
}: CreateNewsletterContainerProps) => {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);

  const handleCreateNewsletter = async (values: FormValues) => {
    setIsLoading(true);
    const token = getCookie("authToken");

    if (!token) {
      console.error("No authentication token found.");
      return;
    }

    const formData = new FormData();
    formData.append("id", uuidv4());
    formData.append("title", values.title);
    formData.append("html", values.html);
    formData.append("categoryId", values.category);

    if (values.emails && values.emails.length > 0) {
      values.emails.forEach((email) => {
        formData.append("recipients[]", email);
      });
    }

    if (values.file) {
      formData.append("file", values.file);
    }

    try {
      const response = await createNewsletter({ formData, token });

      if (response) {
        console.log("Newsletter created successfully:", response);
        router.push("/newsletters");
      }
    } catch (error) {
      console.error("Error creating newsletter:", error);
    }
  };

  return (
    <NewsletterForm
      isLoading={isLoading}
      categories={categories}
      handleSubmit={handleCreateNewsletter}
    />
  );
};
