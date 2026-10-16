package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type DBCommonConfig struct {
    Host        string
    Port        string
    User        string
    Password    string
}

type DBConfig struct {
    DBCommonConfig
    Name        string
}

type Databases struct {
    AuditCkk       *sql.DB
    AuditInternal  *sql.DB
    Certificate     *sql.DB
}

func LoadDBConfigs() map[string]DBConfig {
    common := DBCommonConfig{
        Host:       os.Getenv("DB_HOST"),
        Port:       os.Getenv("DB_PORT"),
        User:       os.Getenv("DB_USER"),
        Password:    os.Getenv("DB_PASSWORD"),
    }

    return map[string]DBConfig{
        "audit_ckk": {
            DBCommonConfig: common,
            Name: os.Getenv("DB_CKK"),
        },
        "audit_internal": {
            DBCommonConfig: common,
            Name: os.Getenv("DB_INTERNAL"),
        },
        "certificate": {
            DBCommonConfig: common,
            Name: os.Getenv("DB_CERT"),
        },
    }
}

func (c DBConfig) DSN() string {
    return fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        c.User,
        c.Password,
        c.Host,
        c.Port,
        c.Name,
    )
}

func InitDatabases(cfg map[string]DBConfig) (*Databases, error) {
    auditCkk, err := sql.Open("mysql", cfg["audit_ckk"].DSN())
    if err != nil {
        log.Fatalf("Connection error: %v", err)
        return nil, err
    }
    if err = auditCkk.Ping(); err != nil {
        log.Fatalf("Database audit_ckk not allowed: %v", err)
    }

    auditInternal, err := sql.Open("mysql", cfg["audit_internal"].DSN())
    if err != nil {
        log.Fatalf("Connection error: %v", err)
        return nil, err
    }
    if err = auditInternal.Ping(); err != nil {
        log.Fatalf("Database audit_internal not allowed: %v", err)
    }

    Certificate, err := sql.Open("mysql", cfg["certificate"].DSN())
    if err != nil {
        log.Fatalf("Connection error: %v", err)
        return nil, err
    }
    if err = Certificate.Ping(); err != nil {
        log.Fatalf("Database certificate not allowed: %v", err)
    }

    log.Println("Successfully connected to databases")
    return &Databases{
        AuditCkk: auditCkk,
        AuditInternal: auditInternal,
        Certificate: Certificate,
    }, nil
}

func (d *Databases) Close() {
    if d.AuditCkk != nil {
        _ = d.AuditCkk.Close()
    }
    if d.AuditInternal != nil {
        _ = d.AuditInternal.Close()
    }
    if d.Certificate != nil {
        _ = d.Certificate.Close()
    }
}


