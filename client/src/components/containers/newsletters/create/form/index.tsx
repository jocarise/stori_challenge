import React from "react";
import { Formik, Form, Field, FieldArray, ErrorMessage } from "formik";
import {
  FormValues,
  validationSchema,
} from "../../../../../schemas/newsletter";

import { Category } from "../../../../../models";
import { CloseIcon } from "../../../../lib/Icons/CloseIcon.component";
import { Button } from "../../../../lib/Button/Button.component";
import { LoaderIcon } from "../../../../lib/Icons/LoaderIcon.component";

interface NewsletterFormProps {
  isLoading: boolean;
  categories: Category[];
  handleSubmit: (values: FormValues) => void;
}

export const NewsletterForm = ({
  isLoading,
  categories,
  handleSubmit,
}: NewsletterFormProps) => {
  return (
    <Formik
      initialValues={{
        title: "",
        html: "",
        file: null,
        category: "",
        emails: [""],
      }}
      validationSchema={validationSchema}
      onSubmit={(values) => {
        handleSubmit(values);
      }}
    >
      {({ setFieldValue, values }) => (
        <Form className="p-6 bg-white rounded-lg shadow-lg">
          <div className="mb-4">
            <label className="block text-gray-700">Newsletter Title:</label>
            <Field
              type="text"
              name="title"
              className="mt-1 block w-full border border-gray-300 rounded p-2"
              disabled={isLoading}
            />
            <ErrorMessage
              name="title"
              component="div"
              className="text-red-600 text-sm"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700">Custom HTML:</label>
            <Field
              disabled={isLoading}
              as="textarea"
              name="html"
              className="mt-1 block w-full border border-gray-300 rounded p-2"
              rows={5}
            />
            <ErrorMessage
              name="html"
              component="div"
              className="text-red-600 text-sm"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700">Attach a File:</label>
            <input
              disabled={isLoading}
              type="file"
              name="file"
              accept=".png, .pdf" // Restrict file types
              onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                const file = event.currentTarget.files?.[0] || null; // Get the file or null
                setFieldValue("file", file); // Set the file value in Formik
              }}
              className="mt-1 block w-full"
            />
            <ErrorMessage
              name="file"
              component="div"
              className="text-red-600 text-sm"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700">Select Category:</label>
            <Field
              disabled={isLoading}
              as="select"
              name="category"
              className="mt-1 block w-full border border-gray-300 rounded p-2"
            >
              <option value="">Select a category</option>
              {categories?.map((c) => (
                <option key={c.id} value={c.id}>
                  {c.title}
                </option>
              ))}
            </Field>
            <ErrorMessage
              name="category"
              component="div"
              className="text-red-600 text-sm"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700">Emails:</label>
            <FieldArray name="emails">
              {({ insert, remove, push }) => (
                <>
                  {values.emails.length > 0 &&
                    values.emails.map((email, index) => (
                      <div key={index}>
                        <div key={index} className="relative">
                          <Field
                            disabled={isLoading}
                            name={`emails.${index}`}
                            className="mt-1 block w-full border border-gray-300 rounded p-2"
                            placeholder="Enter email"
                          />

                          <div className="absolute right-2 top-2">
                            <button
                              disabled={isLoading}
                              type="button"
                              onClick={() => remove(index)}
                              className="bg-red-500 text-white rounded p-1"
                            >
                              <CloseIcon />
                            </button>
                          </div>
                        </div>
                        <ErrorMessage
                          name={`emails.${index}`}
                          component="div"
                          className="text-red-600 text-sm ml-2"
                        />
                      </div>
                    ))}
                  <Button
                    disabled={isLoading}
                    variant="secondary"
                    onClick={() => push("")}
                    className="w-full mt-2"
                  >
                    Add Email
                  </Button>
                </>
              )}
            </FieldArray>
          </div>

          <Button variant="primary" type="submit" disabled={isLoading}>
            {isLoading && <LoaderIcon />} Submit
          </Button>
        </Form>
      )}
    </Formik>
  );
};
