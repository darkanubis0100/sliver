// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the patches on the specified managed node and their
// state relative to the patch baseline being used for the node.
func (c *Client) DescribeInstancePatches(ctx context.Context, params *DescribeInstancePatchesInput, optFns ...func(*Options)) (*DescribeInstancePatchesOutput, error) {
	if params == nil {
		params = &DescribeInstancePatchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstancePatches", params, optFns, c.addOperationDescribeInstancePatchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancePatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstancePatchesInput struct {

	// The ID of the managed node whose patch state information should be retrieved.
	//
	// This member is required.
	InstanceId *string

	// Each element in the array is a structure containing a key-value pair. Supported
	// keys for DescribeInstancePatches include the following:
	//   - Classification Sample values: Security | SecurityUpdates
	//   - KBId Sample values: KB4480056 | java-1.7.0-openjdk.x86_64
	//   - Severity Sample values: Important | Medium | Low
	//   - State Sample values: Installed | InstalledOther | InstalledPendingReboot
	Filters []types.PatchOrchestratorFilter

	// The maximum number of patches to return (per page).
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstancePatchesOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Each entry in the array is a structure containing:
	//   - Title (string)
	//   - KBId (string)
	//   - Classification (string)
	//   - Severity (string)
	//   - State (string, such as "INSTALLED" or "FAILED")
	//   - InstalledTime (DateTime)
	//   - InstalledBy (string)
	Patches []types.PatchComplianceData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstancePatchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstancePatches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstancePatches{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeInstancePatchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstancePatches(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeInstancePatchesAPIClient is a client that implements the
// DescribeInstancePatches operation.
type DescribeInstancePatchesAPIClient interface {
	DescribeInstancePatches(context.Context, *DescribeInstancePatchesInput, ...func(*Options)) (*DescribeInstancePatchesOutput, error)
}

var _ DescribeInstancePatchesAPIClient = (*Client)(nil)

// DescribeInstancePatchesPaginatorOptions is the paginator options for
// DescribeInstancePatches
type DescribeInstancePatchesPaginatorOptions struct {
	// The maximum number of patches to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstancePatchesPaginator is a paginator for DescribeInstancePatches
type DescribeInstancePatchesPaginator struct {
	options   DescribeInstancePatchesPaginatorOptions
	client    DescribeInstancePatchesAPIClient
	params    *DescribeInstancePatchesInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstancePatchesPaginator returns a new
// DescribeInstancePatchesPaginator
func NewDescribeInstancePatchesPaginator(client DescribeInstancePatchesAPIClient, params *DescribeInstancePatchesInput, optFns ...func(*DescribeInstancePatchesPaginatorOptions)) *DescribeInstancePatchesPaginator {
	if params == nil {
		params = &DescribeInstancePatchesInput{}
	}

	options := DescribeInstancePatchesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstancePatchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstancePatchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstancePatches page.
func (p *DescribeInstancePatchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstancePatchesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeInstancePatches(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeInstancePatches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInstancePatches",
	}
}
