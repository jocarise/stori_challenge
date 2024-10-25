export interface Newsletter {
  id: string;
  title?: string; // Default: 'Newsletter'
  attachment?: string;
  html: string;
  createdAt: Date; // autoCreateTime
  updatedAt: Date; // autoUpdateTime
  scheduledDate?: Date; // type: date
  scheduled?: boolean; // Default: false
  categoryId?: number; // nullable
  category?: Category; // optional relation
  recipients?: Recipient[]; // many-to-many relation
}

export interface Category {
  id: number; // primaryKey; autoIncrement
  title: string; // unique
  createdAt: Date; // autoCreateTime
  updatedAt: Date; // autoUpdateTime
  newsletters?: Newsletter[]; // foreignKey: CategoryID
}

export interface Recipient {
  id: string; // primaryKey
  email: string; // unique
  unsuscribeUrl?: string; // type: varchar(255)
  createdAt: Date; // autoCreateTime
  updatedAt: Date; // autoUpdateTime
  newsletters?: Newsletter[]; // many-to-many relation
}
