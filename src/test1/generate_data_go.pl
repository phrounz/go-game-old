#!/usr/bin/perl

use strict;
use warnings;
use File::Basename qw/basename dirname/;
use File::Copy qw/copy/;

my $VERBOSE = 1;

my $SMALL_IMAGES_FOR_ROTATION = 1;
my $PERCENT_SMALL_IMAGES = 40;
my $IMAGEMAGICK_CMD_LINE_UTILITY = 'C:\Users\moi\Desktop\ImageMagick-7.0.8-42-portable-Q16-x64\magick';

open FDW, ">data_go/data.go" || die $!;

print FDW "package data_go\n";
print FDW "\n";
print FDW "func GetBytesFromFilename(filename string) []byte {\n";
print FDW '  if len(filename)>=2 && filename[0:2]!="./" { filename = "./"+filename }'."\n";
print FDW "  switch (filename) {\n";

sub doFolder {
  my ($root, $folder) = @_;
  foreach my $file (glob("$folder/*")) {
    if (-d $file) {
      doFolder($root, $file);
    } elsif ($file =~ /^$root\/(.*)\.png$/) {
      my $outvar = $1."_png";
      $outvar =~ s/-/_/g;
      $outvar =~ s/\//__/g;
      my $outfile = "./data_go/$outvar.go";
      mkdir dirname($outfile) unless (-d dirname($outfile));
      # my $bn = basename($file, ".png");
      #print FDW '//go:generate ../file2byteslice -package='.$package_name.' -input='.$file.' -output='.$outfile.' -var='.$bn.'_png'."\n"

      my $tmpfile = $file;
      if ($SMALL_IMAGES_FOR_ROTATION && ($file=~/\/small\//)) {
        $tmpfile = "tmp_file_image_magick~";
        copy($file, $tmpfile);
        my $cmd_resize ="$IMAGEMAGICK_CMD_LINE_UTILITY mogrify -resize $PERCENT_SMALL_IMAGES% $tmpfile";
        print "=> $cmd_resize\n" if $VERBOSE;
        print `$cmd_resize`;
      }

      my $cmd = "./file2byteslice -package=data_go -input=$tmpfile -output=$outfile -var=$outvar ";
      print "=> $cmd\n" if $VERBOSE;
      print `$cmd`;

      print FDW "  case \"$file\": return $outvar\n";

    } elsif ($file =~ /\~$/) {
      # ignore
    } else {
      die $file;
    }
  }
}
doFolder("./data", "./data");

#print FDW "//go:generate gofmt -s -w .\n";
print FDW "  default: panic(\"Could not find: \"+filename)\n";
print FDW "  }\n";
print FDW "  return []byte{}\n";
print FDW "}\n";

close FDW;
