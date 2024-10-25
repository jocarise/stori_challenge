import { Button } from "../Button/Button.component";

interface NewsletterCard {
  id: string;
  date: string;
  dateEnabled: boolean;
  title: string;
  badgeTitle: string;
  mainActionCopy: string;
  secondaryActionCopy: string;
  handlerMainAction: (id: string) => void;
  handleSecondaryAction: (id: string) => void;
  mainIcon?: React.ReactNode;
  secondaryIcon?: React.ReactNode;
}

export const NewsletterCard = (props: NewsletterCard) => {
  const {
    id,
    dateEnabled,
    date,
    title,
    badgeTitle,
    mainActionCopy,
    secondaryActionCopy,
    handlerMainAction,
    handleSecondaryAction,
    mainIcon,
    secondaryIcon,
  } = props;

  return (
    <div
      key={id}
      className="max-w-4xl px-10 my-4 py-6 bg-white rounded-lg shadow-md"
    >
      <div className="flex justify-end">
        {dateEnabled && (
          <span className="font-light text-gray-600 mr-4">{date}</span>
        )}
        <span className="bg-gray-100 text-gray-800 text-sm font-medium me-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">
          {badgeTitle}
        </span>
      </div>
      <div className="mt-2">
        <p className="text-2xl text-gray-700 font-bold">{title}</p>
      </div>
      <div className="flex justify-end mt-8">
        <Button
          variant="secondary"
          onClick={() => {
            handleSecondaryAction(id);
          }}
          className="mr-4"
        >
          {secondaryIcon && secondaryIcon}
          {secondaryActionCopy}
        </Button>
        <div>
          <Button
            variant="primary"
            onClick={() => {
              handlerMainAction(id);
            }}
          >
            {mainIcon && mainIcon} {mainActionCopy}
          </Button>
        </div>
      </div>
    </div>
  );
};
