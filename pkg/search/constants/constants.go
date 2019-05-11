package constants

const (
	DaemonsetHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE"
	DaemonsetHeaderWide  = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	DeploymentHeader     = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tAGE"
	DeploymentHeaderWide = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tAGE\tCONTAINERS\tIMAGES"
	HpaHeader            = "NAMESPACE\tNAME\tREFERENCE\tTARGETS\tMINPODS\tMAXPODS\tREPLICAS\tAGE"
	NodeHeader           = "NAME\tSTATUS\tROLES\tAGE\tVERSION"
	NodeHeaderWide       = "NAME\tSTATUS\tROLES\tAGE\tVERSION\tINTERNAL-IP\tEXTERNAL-IP\tOS-IMAGE\tKERNEL-VERSION\tCONTAINER-RUNTIME"
	PodHeader            = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTART\tAGE"
	PodHeaderWide        = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTART\tAGE\tIP\tNODENAME"

	DaemonsetRowTemplate      = "%s\t%s\t%d\t%d\t%d\t%d\t%s\t%d%s"
	DaemonsetRowTemplateWide  = "%s\t%s\t%d\t%d\t%d\t%d\t%s\t%d%s\t%s\t%s\t%s"
	DeploymentRowTemplate     = "%s\t%s\t%d\t%d\t%d\t%d\t%d%s"
	DeploymentRowTemplateWide = "%s\t%s\t%d\t%d\t%d\t%d\t%d%s\t%s\t%s"
	HpaRowTemplate            = "%s\t%s\t%s/%s\t%d%%/%d%%\t%d\t%d\t%d\t%d%s"
	NodeRowTemplate           = "%s\t%s\t%s\t%d%s\t%s"
	NodeRowTemplateWide       = "%s\t%s\t%s\t%d%s\t%s\t%s\t%s\t%s\t%s\t%s"
	PodRowTemplate            = "%s\t%s\t%d/%d\t%s\t%d\t%d%s"
	PodRowTemplateWide        = "%s\t%s\t%d/%d\t%s\t%d\t%d%s\t%s\t%s"
)