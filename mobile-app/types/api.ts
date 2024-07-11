export interface Baby extends NewBaby {
  id: number;
}

export interface NewBaby {
  userID: number;
  name: string;
  birthDate: string;
}

export interface Event extends NewEvent {
  id: number;
  babyID: number;
}

export interface NewEvent {
  eventType: string;
  startTime: string;
  endTime?: string;
  notes?: string;
}
