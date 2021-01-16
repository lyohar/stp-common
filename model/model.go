package model

const ExportSkipCommand = "SKIP"

type Export struct {
	Command              string `json:"command" binding:"required"`
	TableSchema          string `json:"table_schema" binding:"required"`
	TableName            string `json:"table_name" binding:"required"`
	ColumnsList          string `json:"columns_list" binding:"required"`
	KeysList             string `json:"key_fields" binding:"required"`
	OrderColumnName      string `json:"order_column_name" binding:"required"`
	OrderColumnFromValue string `json:"order_column_value" binding:"required"`
	TimestampColumnName  string `json:"timestamp_column_name" binding:"required"`
	OperationColumnName  string `json:"operation_column_name" binding:"required"`
	TopicName            string `json:"topic_name" binding:"required"`
}

type ExportStatus struct {
	TableSchema        string `json:"table_schema" binding:"required"`
	TableName          string `json:"table_name" binding:"required"`
	OrderColumnToValue string `json:"order_column_value" binding:"required"`
}
