import { useState } from "react";
import { ErrorMessage, Field, Form, Formik } from "formik";
import { getCookie } from "cookies-next";
import { Newsletter } from "../../../models";
import { formatDate } from "../../../utils";
import { validationSchema } from "../../../schemas/email";
import { addRecipient } from "../../../adapters/addRecipient";
import { sendEmails } from "../../../adapters/sendEmails";
import { ModalBody, ModalContainer, ModalFooterActions } from "../../lib/Modal";
import { NewsletterCard } from "../../lib/NewsletterCard/NewsletterCard.component";

interface NewsletterContainerProps {
  newsletters: Newsletter[];
}

export const NewsletterContainer = ({
  newsletters,
}: NewsletterContainerProps) => {
  const [loading, setLoading] = useState(false);
  const [isSendEmailModalOpen, setIsisSendEmailModalOpen] = useState(false);
  const [isCreateRecipientModalOpen, setIsCreateRecipientModalOpen] =
    useState(false);
  const [selectedNewsletter, setSelectedNewsletter] = useState("");

  const handleOpenSendEmailModal = (newsletterId: string) => {
    setIsisSendEmailModalOpen(true);
    setSelectedNewsletter(newsletterId);
  };

  const handleOpenRecipientModal = (newsletterId: string) => {
    setIsCreateRecipientModalOpen(true);
    setSelectedNewsletter(newsletterId);
  };

  const handleToggleCreateRecipientModal = () =>
    setIsCreateRecipientModalOpen((prev) => !prev);
  const handleToggleSendEmailModal = () =>
    setIsisSendEmailModalOpen((prev) => !prev);

  const handleAddRecipient = async (email: string) => {
    const token = getCookie("authToken");

    if (token) {
      const response = await addRecipient({
        email,
        newsletterId: selectedNewsletter,
        token,
      });
      if (!response) {
        console.log(response);
      }
    }

    setIsCreateRecipientModalOpen(false);
  };

  const handleSendEmails = async () => {
    setLoading(true);
    const token = getCookie("authToken");

    if (token) {
      const newsletterId = selectedNewsletter;
      const status = await sendEmails(token, newsletterId);

      if (status != 200) {
        console.error("error sending emails");
      }

      setLoading(false);
      setIsisSendEmailModalOpen(false);
    }
  };

  return (
    <>
      {newsletters?.map((n) => {
        return (
          <NewsletterCard
            id={n.id}
            key={n.id}
            dateEnabled={n?.scheduled ? n.scheduled : false}
            date={formatDate(n?.scheduledDate ? n?.scheduledDate : new Date())}
            title={n?.title ? n.title : ""}
            badgeTitle={n?.category?.title ? n.category.title : ""}
            mainActionCopy={"Send Newsletter"}
            secondaryActionCopy={"Add recipient"}
            handlerMainAction={handleOpenSendEmailModal}
            handleSecondaryAction={handleOpenRecipientModal}
          />
        );
      })}

      {isSendEmailModalOpen && (
        <ModalContainer showModal={isSendEmailModalOpen}>
          <ModalBody>
            Are you sure to send newsletter to all recipients?
          </ModalBody>
          <ModalFooterActions
            isLoading={loading}
            primaryButtonCopy={"Confirm"}
            secondaryButtonCopy={"Close"}
            handlePrimaryAction={handleSendEmails}
            handleSecondaryAction={handleToggleSendEmailModal}
            className="pb-3.5 pr-3.5"
          />
        </ModalContainer>
      )}

      {isCreateRecipientModalOpen && (
        <ModalContainer showModal={isCreateRecipientModalOpen}>
          <Formik
            initialValues={{ email: "" }}
            validationSchema={validationSchema}
            onSubmit={({ email }) => {
              handleAddRecipient(email);
            }}
          >
            {() => (
              <Form className="mx-auto p-6 bg-white rounded-lg shadow-lg w">
                <div className="mb-4">
                  <label className="block text-gray-700">Email:</label>
                  <Field
                    type="email"
                    name="email"
                    className="mt-1 block w-full border border-gray-300 rounded p-2 focus:outline-none focus:ring focus:ring-indigo-500"
                  />
                  <ErrorMessage
                    name="email"
                    component="div"
                    className="text-red-600 text-sm mt-1"
                  />
                </div>

                <ModalFooterActions
                  isLoading={false}
                  primaryButtonCopy={"Add"}
                  secondaryButtonCopy={"Close"}
                  handlePrimaryAction={() => {}}
                  handleSecondaryAction={handleToggleCreateRecipientModal}
                />
              </Form>
            )}
          </Formik>
        </ModalContainer>
      )}
    </>
  );
};
