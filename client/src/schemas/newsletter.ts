import * as Yup from "yup";

export interface FormValues {
  title: string;
  html: string;
  file: File | null;
  category: string;
  emails: string[];
}

export const validationSchema = Yup.object<FormValues>().shape({
  title: Yup.string().required("Title is required"),
  html: Yup.string(),
  file: Yup.mixed()
    .required("File is required")
    .test("fileType", "Only PNG and PDF files are allowed", (value) => {
      return (
        value &&
        value instanceof File && // Check if value is a File
        (value.type === "image/png" || value.type === "application/pdf")
      );
    })
    .test("fileSize", "File size must be 5 MB or less", (value) => {
      return (
        value &&
        value instanceof File && // Check if value is a File
        value.size <= 5 * 1024 * 1024 // 5 MB in bytes
      );
    }),
  category: Yup.string().required("Category is required"),
  emails: Yup.array()
    .of(Yup.string().email("Invalid email format"))
    .required("At least one email is required"),
});
