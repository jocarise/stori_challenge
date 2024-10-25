import { Button } from "../Button/Button.component";
import { LoaderIcon } from "../Icons/LoaderIcon.component";

interface ModalFooterProps {
  isLoading?: boolean;
  primaryButtonCopy: string;
  handlePrimaryAction: () => void;
}

export const ModalFooterAction = ({
  isLoading = false,
  primaryButtonCopy,
  handlePrimaryAction,
}: ModalFooterProps) => {
  return (
    <>
      <div className="flex items-center justify-end rounded-b">
        <Button
          type="submit"
          variant="primary"
          disabled={isLoading}
          onClick={handlePrimaryAction}
        >
          {isLoading && <LoaderIcon />} {primaryButtonCopy}
        </Button>
      </div>
    </>
  );
};
