export interface Domain {
  domain: string;
  serversChanged: boolean;
  sslGrade: string;
  previousSslGrade: string;
  logo: string;
  title: string;
  isDown: boolean;
  servers: Server[];
}

export interface DomainResponse {
  items: Domain[];
}

export interface Server {
  address: string;
  sslGrade: string;
  country: string;
  owner: string;
}
