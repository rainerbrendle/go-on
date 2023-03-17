/*
 * SCHEMA PR2
 * 
 * JOURNAL
 */

DROP SCHEMA PR2 CASCADE;
CREATE SCHEMA PR2;


CREATE TABLE IF NOT EXISTS PR2.Journal (
   ClassName VARCHAR NOT NULL,      /* 'Actor' Class Name */
   Recipient VARCHAR NOT NULL,      /* Recipient Location (shard) */
   Sender VARCHAR NOT NULL,         /* Sender location */
   SenderSequence BigInt NOT NULL,  /* Sequence message outbox queue */
   MessageSequence BigInt NOT NULL, /* Sequence (Offset) */
   Reason VARCHAR,                  /* Reason (usually 'account #) */
   DateTime  VARCHAR,               /* Physical Recording Time */
   Epoch VARCHAR,                   /* Epoche (for fail-over) */
   bData BYTEA,                     /* Binary Data (optional) */
   cData VARCHAR,                   /* (Unicode) String Data */
   jData JSONB,                     /* JSON payload */
   MajorProtocolVersion BigInt,     /* Protocol Version */
   MinorProtocolVersion BigInt,     /* (Minor) Protocol Version */
   Signature VARCHAR,               /* Signature */

   UNIQUE(ClassName, Recipient, Sender, SenderSequence )
 );



